package parser

import (
    "testing"
    "os"
)

func TestGetReplayHeader(t *testing.T) {
    /*
     this test file has 136 bytes.
     12 bytes for the signature, 2 bytes for the metadata size
     then padded until 0x1A
     */
    const test_file_path = "test_files/METADATA_ONLY_test_replay_20220127_1829_Klantsmurfen_RO_vs_Nibnab_JA.ggr"

    file, _ := os.Open(test_file_path)

    defer file.Close()

    header, err := GetReplayHeader(file)
    if err != nil {
        t.Fatalf("%q", err)
    }

    if header.Metadata_size != 110 {
        t.Fatalf("This test file's metadata size should be 110 bytes.")
    }
}

