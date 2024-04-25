package api

import (
	"time"

	jsoniter "github.com/json-iterator/go"
	jsontime "github.com/liamylian/jsontime/v2/v2"
)

var json = jsontime.ConfigWithCustomTimeFormat

func init() {
	// Antelope Api does not specify timezone in timestamps (they are always UTC tho).
	jsontime.SetDefaultTimeFormat("2006-01-02T15:04:05", time.UTC)
}

func Json() jsoniter.API {
	return json
}

func customJsonUnmarshal(data []byte, v interface{}) error {
	// Empty data is valid.
	if len(data) < 1 {
		return nil
	}
	return json.Unmarshal(data, v)
}
