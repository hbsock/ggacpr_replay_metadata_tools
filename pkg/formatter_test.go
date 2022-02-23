package metadata

import (
	"reflect"
	"testing"
	"time"
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
		Date:    date,
		Content: content,
	}

}

func TestMetaDataToDate(t *testing.T) {
	metadata := _GetTestData()

	v_time := metadata.GetDate()
	t.Logf("%v", v_time)

	// 2022-01-27 18:29:58 -0500
	v_timezone := time.FixedZone("EST", -18000)
	v_expected_time := time.Date(2022, 1, 27, 18, 29, 58, 0, v_timezone)

	if !v_time.Equal(v_expected_time) {
		t.Fatalf("Time %v did not match the expected time %v", v_time, v_expected_time)
	}
}

func TestMetaDataToStringSlice(t *testing.T) {
	metadata := _GetTestData()

	ss := metadata.ToStringSlice()

	expected_ss := []string{
		"2022-01-27 18:29:58 -0500 UTC-5",
		"76561198008085514",
		"76561198011058687",
		"Klantsmurfen",
		"Nibnab",
		"RO",
		"JA",
		"0",
		"1",
		"0",
		"2",
		"0",
		"0",
		"124",
		"6810",
		"4",
		"0",
		"5",
		"3",
		"1",
	}
	if !reflect.DeepEqual(ss, expected_ss) {

		t.Fatalf("Expected ss %#v did not match the actual ss %#v", expected_ss, ss)
	}
}

func TestMetaDataToHeaders(t *testing.T) {
	headers := GetReplayMetadataHeaders()
	expected_len := 20
	if len(headers) != expected_len {
		t.Fatalf("Expected length of headers is %v", expected_len)
	}
}
