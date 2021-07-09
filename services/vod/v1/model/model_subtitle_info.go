package model

import (
	"encoding/json"

	"strings"
)

type SubtitleInfo struct {
	// 字幕文件的下载地址

	Url *string `json:"url,omitempty"`
	// 字幕文件id

	Id *int32 `json:"id,omitempty"`
	// 字幕文件类型

	Type *string `json:"type,omitempty"`
	// 字幕文件语言种类

	Language *string `json:"language,omitempty"`
}

func (o SubtitleInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SubtitleInfo struct{}"
	}

	return strings.Join([]string{"SubtitleInfo", string(data)}, " ")
}
