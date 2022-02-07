package parser

import (
    "io"
    "encoding/binary"
    "errors"
)

type ReplayHeader struct {
    Signature [12]byte
    Metadata_size uint16
    _ [12]byte //To skip to the 0x1A byte where the metadata starts
}

type ReplayRecordingDate struct {
    Year uint16
    Month uint8
    Day uint8
    Hour uint8
    Minute uint8
    Second uint8
    _ [1]byte //To skip to the 0x22 byte
}

type ReplayMetaDataContent struct {
    Player1SteamID uint64
    Player2SteamID uint64
    Player1NameUTF8 [32]byte
    Player2NameUTF8 [32]byte

}


type ReplayMetaData struct {
    Header ReplayHeader
    Date ReplayRecordingDate
    Content ReplayMetaDataContent
}


func GetReplayHeader(r io.Reader) (ReplayHeader, error) {
    header := ReplayHeader{}
    err := binary.Read(r, binary.LittleEndian, &header )
    if err != nil {
        return header, err
    }

    SIGNATURE := [12]byte{0x47, 0x47, 0x52, 0x02, 0x51, 0xad, 0xee, 0x77, 0x45, 0xd7, 0x48, 0xcd}
    if header.Signature != SIGNATURE {
        return header, errors.New("The file format signature is not a valid GG AC+R replay file")
    } else if header.Metadata_size != 110 {
        return header, errors.New("The metadata size is not 110, this replay file is likely not from the latest version of GG AC+R. Please watch this replay file until the end on the latest game version to update the metadata.")
    }

    return header, nil
}


func GetReplayRecordingDate(r io.Reader) (ReplayRecordingDate, error) {
    recording_date := ReplayRecordingDate{}
    err := binary.Read(r, binary.LittleEndian, &recording_date )
    if err != nil {
        return recording_date, err
    }

    return recording_date, nil
}

func GetReplayMetaData(r io.Reader) (ReplayMetaData, error) {
    metadata := ReplayMetaData{}
    err := binary.Read(r, binary.LittleEndian, &metadata)
    if err != nil {
        return metadata, err
    }

    return metadata, nil
}

