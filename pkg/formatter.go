package metadata

import (
	"bytes"
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


func (md ReplayMetaData) ToStringSlice() []string {
	var s []string

	// Need the .rstrip() to remove the null \x00 values from the remaining 32 bytes of P1 or P2 names.
	s = append(s, md.Content.GetP1NameStr() )
	s = append(s, md.Content.GetP2NameStr() )

	return s
}
