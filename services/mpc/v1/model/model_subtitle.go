package model

import (
	"encoding/json"

	"strings"
)

type Subtitle struct {
	Input *ObsObjInfo `json:"input,omitempty"`
	// 多字幕文件地址。
	Inputs *[]MulInputFileInfo `json:"inputs,omitempty"`
	// 字幕类型
	SubtitleType *int32 `json:"subtitle_type,omitempty"`
}

func (o Subtitle) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Subtitle struct{}"
	}

	return strings.Join([]string{"Subtitle", string(data)}, " ")
}
