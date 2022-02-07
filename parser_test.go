package parser

import (
    "testing"
    "os"
)

const TEST_FILE_PATH = "test_files/METADATA_ONLY_test_replay_20220127_1829_Klantsmurfen_RO_vs_Nibnab_JA.ggr"

func TestGetReplayHeader(t *testing.T) {
    /*
     this test file has 136 bytes.
     12 bytes for the signature, 2 bytes for the metadata size
     then padded until 0x1A
     */

    file, _ := os.Open(TEST_FILE_PATH)
    defer file.Close()

    _, err := GetReplayHeader(file)
    if err != nil {
        t.Fatalf("%q", err)
    }
}


func TestGetReplayRecordingDate(t *testing.T) {
    file, _ := os.Open(TEST_FILE_PATH)
    defer file.Close()

    GetReplayHeader(file)

    recording_date, err := GetReplayRecordingDate(file)
    if err != nil {
        t.Fatalf("%q", err)
    }

    expected_date := ReplayRecordingDate{
        Year: 2022,
        Month: 1,
        Day: 27,
        Hour: 18,
        Minute: 29,
        Second: 58,
    }

    if recording_date != expected_date {
        t.Logf("Expected date was %+q but read %+q", expected_date, recording_date)
        t.Fatalf("Expected date was %+v but read %+v", expected_date, recording_date)
    }
}


func TestGetReplayMetaData(t *testing.T) {
    file, _ := os.Open(TEST_FILE_PATH)
    defer file.Close()

    metadata, err := GetReplayMetaData(file)
    if err != nil {
        t.Fatalf("%q", err)
    }

    expected_content := ReplayMetaDataContent{
        Player1SteamID: 76561198008085514,
        Player2SteamID: 76561198011058687,
    }
    copy(expected_content.Player1NameUTF8[:], []byte("Klantsmurfen"))
    copy(expected_content.Player2NameUTF8[:], []byte("Nibnab"))

    t.Log(metadata)

    if expected_content != metadata.Content {
        t.Fatalf("Expected content %v did not match the actual content %v", expected_content, metadata.Content)
    }
}
