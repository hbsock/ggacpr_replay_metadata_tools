package metadata

import (
	"testing"
	"reflect"
)

func _GetTestData() ReplayMetaData {

	date := ReplayRecordingDate{
		Year:   2022,
		Month:  1,
		Day:    27,
		Hour:   18,
		Minute: 29,
		Second: 58,
	}

	content := ReplayMetaDataContent{
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
	copy(content.P1NameUTF8[:], []byte("Klantsmurfen"))
	copy(content.P2NameUTF8[:], []byte("Nibnab"))

	return ReplayMetaData{
		Date: date,
		Content: content,
	}

}


func TestMetaDataToStringSlice(t *testing.T) {
	metadata := _GetTestData()

	ss := metadata.ToStringSlice()

	expected_ss := []string{ "Klantsmurfen", "Nibnab"}
	if !reflect.DeepEqual(ss, expected_ss) {

		t.Fatalf("Expected ss %#v did not match the actual ss %#v", expected_ss, ss)
	}
}
