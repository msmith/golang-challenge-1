package drum

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"os"
)

// SPLICE file format:
//
// File header
//       6 bytes - file signature ("SPLICE")
//       8 bytes - length of the pattern header + track data (uint64 big-endian)
//
// Pattern header
//      32 bytes - hardware version string (null-terminated) & padding
//       4 bytes - tempo (float32 little-endian)
//
// Track data (repeated)
//       4 bytes - track ID (uint32 little-endian)
//       1 byte  - track name length (uint8)
//   0-255 bytes - track name
//      16 bytes - track steps (1 = play sound, 0 = silence)

var fileSignature = []byte("SPLICE")

// DecodeFile will read a file and return the Pattern
func DecodeFile(path string) (*Pattern, error) {
	// open file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// decode the file's contents
	return NewDecoder().Decode(f)
}

func NewDecoder() *decoder {
	return &decoder{
		h: header{
			Signature: make([]byte, 6),
			HwVers:    make([]byte, 32),
		},
	}
}

type decoder struct {
	h  header
	td trackDecoder
}

// Decode will read from a reader and return the Pattern
func (d *decoder) Decode(reader io.Reader) (*Pattern, error) {
	// Decode file header
	if _, err := io.ReadFull(reader, d.h.Signature[:]); err != nil {
		return nil, err
	}

	// Ensure that signature matches the expected value
	if !bytes.Equal(d.h.Signature[:], fileSignature) {
		return nil, errors.New("SPLICE file signature not found")
	}

	if err := binary.Read(reader, binary.BigEndian, &d.h.Length); err != nil {
		return nil, err
	}

	// Create a limited reader to read more more than h.Length bytes
	dataReader := io.LimitReader(reader, int64(d.h.Length))

	// HW Version
	if _, err := io.ReadFull(dataReader, d.h.HwVers); err != nil {
		return nil, err
	}

	// Tempo
	if err := binary.Read(dataReader, binary.LittleEndian, &d.h.Tempo); err != nil {
		return nil, err
	}

	// Pattern
	pattern := &Pattern{
		Version: string(d.h.HwVers[:bytes.IndexByte(d.h.HwVers, 0)]),
		Tempo:   d.h.Tempo,
	}

	// Decode all the Tracks
	for {
		track, err := d.td.decode(dataReader)
		if err == io.EOF {
			// no more tracks, exit the loop
			break
		}
		if err != nil {
			return nil, err
		}
		pattern.AddTrack(*track)
	}

	return pattern, nil
}

type header struct {
	Signature []byte
	Length    uint64
	HwVers    []byte
	Tempo     float32
}

type trackDecoder struct {
	len   uint8
	tmp   [256]byte // big enough for the maximum name length
	track Track
}

func (t *trackDecoder) decode(reader io.Reader) (*Track, error) {
	// ID: 4 bytes
	if err := binary.Read(reader, binary.LittleEndian, &t.track.ID); err != nil {
		return nil, err
	}

	// nameLength: 1 byte
	if err := binary.Read(reader, binary.LittleEndian, &t.len); err != nil {
		return nil, err
	}

	// Name: nameLength bytes
	nameBytes := t.tmp[0:t.len]
	if _, err := io.ReadFull(reader, nameBytes); err != nil {
		return nil, err
	}
	t.track.Name = string(nameBytes)

	// steps: 16 bytes
	steps := t.tmp[0:16]
	if _, err := io.ReadFull(reader, steps); err != nil {
		return nil, err
	}

	// convert step bytes into booleans
	for i, val := range steps {
		t.track.Steps[i] = (val > 0)
	}

	return &(t.track), nil
}
