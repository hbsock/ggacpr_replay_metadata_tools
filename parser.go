package parser

import (
    "io"
    "encoding/binary"
    "errors"
)

type ReplayHeader struct {
    Signature [12]byte
    Metadata_size uint16
    _ [14]byte
}


func GetReplayHeader(r io.Reader) (ReplayHeader, error) {
    header := ReplayHeader{}
    binary.Read(r, binary.LittleEndian, &header )

    SIGNATURE := [12]byte{0x47, 0x47, 0x52, 0x02, 0x51, 0xad, 0xee, 0x77, 0x45, 0xd7, 0x48, 0xcd}
    if header.Signature != SIGNATURE {
        return header, errors.New("The file format signature is not a valid GG AC+R replay file")
    }

    return header, nil
}


