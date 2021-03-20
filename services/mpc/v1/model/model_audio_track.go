package model

import (
	"encoding/json"

	"strings"
)

type AudioTrack struct {
	// 音轨选取方式。 - 0：默认选取 - 1：手动选择

	Type *int32 `json:"type,omitempty"`
	// 选取左声道所在的音轨编号。

	Left *int32 `json:"left,omitempty"`
	// 选取右声道所在的音轨编号。

	Right *int32 `json:"right,omitempty"`
}

func (o AudioTrack) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AudioTrack struct{}"
	}

	return strings.Join([]string{"AudioTrack", string(data)}, " ")
}
