package metadata

import (
	"github.com/hbsock/ggacpr_replay_metadata_tools/internal/projectpath"
	"os"
	"path/filepath"
	"testing"
)

var (
	TEST_FILE_PATH = filepath.Join(projectpath.Root, "test_files/METADATA_ONLY_test_replay_20220127_1829_Klantsmurfen_RO_vs_Nibnab_JA.ggr")
)

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
		Year:   2022,
		Month:  1,
		Day:    27,
		Hour:   18,
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
		P1SteamID:             76561198008085514,
		P2SteamID:             76561198011058687,
		P1Character:           RO,
		P2Character:           JA,
		ExtraFlag:             0,
		SingleOrTeamFlag:      1,
		IsPlusRFlag:           0,
		TimezoneBiasInSec:     18000,
		P1RoundsWon:           2,
		P2RoundsWon:           0,
		ProblemBitmask:        0,
		Ping:                  124,
		MatchDurationInFrames: 6810,
		P1Score:               4,
		P2Score:               0,
		P1Rank:                5,
		P2Rank:                3,
		WinnerSide:            1,
	}
	copy(expected_content.P1NameUTF8[:], []byte("Klantsmurfen"))
	copy(expected_content.P2NameUTF8[:], []byte("Nibnab"))

	//t.Logf("Sol value is %d", SO)

	if expected_content != metadata.Content {
		t.Fatalf("Expected content %v did not match the actual content %v", expected_content, metadata.Content)
	}
}
