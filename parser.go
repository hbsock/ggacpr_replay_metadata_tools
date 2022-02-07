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

type Character uint8
const (
    SO Character = iota + 1
    KY
    MA
    MI
    AX
    PO
    CH
    ED
    BA
    FA
    TE
    JA
    AN
    JO
    VE
    DI
    SL
    IN
    ZA
    BR
    RO
    AB
    OS
    KL
    JU
)

type ReplayMetaDataContent struct {
    P1SteamID uint64
    P2SteamID uint64
    P1NameUTF8 [32]byte
    P2NameUTF8 [32]byte
    P1Character Character
    P2Character Character
    ExtraFlag uint8 // value is 1 if extra options were modified, or EX/SP/GG characters were used, or KL/JU in the original AC, 0 otherwise
    SingleOrTeamFlag uint8 // 1 = single, 2 = team
    IsPlusRFlag uint8 // 0 = +R, 1 = AC
    TimezoneBiasInSec int32 //
    P1RoundsWon uint8
    P2RoundsWon uint8
    ProblemBitmask uint8 // 0b1 = unfinished match 0b10 = disconnect 0b100 = desync
    Ping uint8
    MatchDurationInFrames uint32
    P1Score uint8
    P2Score uint8
    P1Rank uint8
    P2Rank uint8
    WinnerSide uint8 // 1 = P1, 2 = P2, 3 = draw
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

