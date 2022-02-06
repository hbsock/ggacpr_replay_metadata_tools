package parser

import (
    "io"
    "encoding/binary"
    "errors"
)

type ReplayHeader struct {
    Signature [12]byte
    Metadata_size uint16
    _ [14]byte //To skip to the 0x1A byte where the metadata starts
}


func GetReplayHeader(r io.Reader) (ReplayHeader, error) {
    header := ReplayHeader{}
    binary.Read(r, binary.LittleEndian, &header )

    SIGNATURE := [12]byte{0x47, 0x47, 0x52, 0x02, 0x51, 0xad, 0xee, 0x77, 0x45, 0xd7, 0x48, 0xcd}
    if header.Signature != SIGNATURE {
        return header, errors.New("The file format signature is not a valid GG AC+R replay file")
    } else if header.Metadata_size != 110 {
        return header, errors.New("The metadata size is not 110, this replay file is likely not from the latest version of GG AC+R. Please watch this replay file until the end on the latest game version to update the metadata.")
    }

    return header, nil
}

