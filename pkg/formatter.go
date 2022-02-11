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
	timezone_bias_in_secs := int(-md.Content.TimezoneBiasInSec)
	timezone := time.FixedZone(
		"UTC" + strconv.Itoa(timezone_bias_in_secs / (60 * 60) ),
		timezone_bias_in_secs,
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

func (e Character) String() string {
	switch e {
	case SO:
		return "SO"
	case KY:
		return "KY"
	case MA:
		return "MA"
	case MI:
		return "MI"
	case AX:
		return "AX"
	case PO:
		return "PO"
	case CH:
		return "CH"
	case ED:
		return "ED"
	case BA:
		return "BA"
	case FA:
		return "FA"
	case TE:
		return "TE"
	case JA:
		return "JA"
	case AN:
		return "AN"
	case JO:
		return "JO"
	case VE:
		return "VE"
	case DI:
		return "DI"
	case SL:
		return "SL"
	case IN:
		return "IN"
	case ZA:
		return "ZA"
	case BR:
		return "BR"
	case RO:
		return "RO"
	case AB:
		return "AB"
	case OS:
		return "OS"
	case KL:
		return "KL"
	case JU:
		return "JU"
	default:
		return ""
	}
}


func (md ReplayMetaData) ToStringSlice() []string {
	var s []string

	s = append(s, md.GetDate().String())
	s = append(s, strconv.FormatUint(md.Content.P1SteamID, 10) )
	s = append(s, strconv.FormatUint(md.Content.P2SteamID, 10) )
	s = append(s, md.Content.GetP1NameStr())
	s = append(s, md.Content.GetP2NameStr())
	s = append(s, md.Content.P1Character.String())
	s = append(s, md.Content.P2Character.String())
	s = append(s, strconv.FormatUint( uint64(md.Content.ExtraFlag), 10) )
	s = append(s, strconv.FormatUint( uint64(md.Content.SingleOrTeamFlag), 10) )
	s = append(s, strconv.FormatUint( uint64(md.Content.IsPlusRFlag), 10) )
	s = append(s, strconv.FormatUint( uint64(md.Content.P1RoundsWon), 10) )
	s = append(s, strconv.FormatUint( uint64(md.Content.P2RoundsWon), 10) )
	s = append(s, strconv.FormatUint( uint64(md.Content.ProblemBitmask), 10) )
	s = append(s, strconv.FormatUint( uint64(md.Content.Ping), 10) )
	s = append(s, strconv.FormatUint( uint64(md.Content.MatchDurationInFrames), 10) )
	s = append(s, strconv.FormatUint( uint64(md.Content.P1Score), 10) )
	s = append(s, strconv.FormatUint( uint64(md.Content.P2Score), 10) )
	s = append(s, strconv.FormatUint( uint64(md.Content.P1Rank), 10) )
	s = append(s, strconv.FormatUint( uint64(md.Content.P2Rank), 10) )
	s = append(s, strconv.FormatUint( uint64(md.Content.WinnerSide), 10) )


	return s
}
