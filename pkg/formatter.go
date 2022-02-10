package metadata

import (
	"bytes"
	"time"
	"strconv"
)

func _GetStringFromBytes(b []byte) string {
	b = bytes.Trim(b, "\x00")
	return string(b)
}

func (c ReplayMetaDataContent) GetP1NameStr() string {
	return _GetStringFromBytes(c.P1NameUTF8[:])
}

func (c ReplayMetaDataContent) GetP2NameStr() string {
	return _GetStringFromBytes(c.P2NameUTF8[:])
}

func (md ReplayMetaData) GetDate() time.Time {
	// taking negative Timezone bias since my replays are from EST which is UTC-5 hours, but the parsed value seems to be
	// positive 5. So I'm guessing the signs are reversed?
	timezone_bias_in_hours := int(-md.Content.TimezoneBiasInSec) / (60 * 60)
	timezone := time.FixedZone(
		"UTC" + strconv.Itoa(timezone_bias_in_hours),
		int(-md.Content.TimezoneBiasInSec),
	)


	return time.Date(
		int(md.Date.Year),
		time.Month(md.Date.Month),
		int(md.Date.Day),
		int(md.Date.Hour),
		int(md.Date.Minute),
		int(md.Date.Second),
		0,
		timezone,
	)
}

func (md ReplayMetaData) ToStringSlice() []string {
	var s []string

	s = append(s, md.GetDate().String())
	s = append(s, md.Content.GetP1NameStr())
	s = append(s, md.Content.GetP2NameStr())


	return s
}
