package model

import (
	"encoding/json"

	"strings"
)

type VodSampleData struct {
	// 存储空间，单位：GB <br/>

	Storage *float32 `json:"storage,omitempty"`
	// 转码时长，单位：秒 <br/>

	Transcode *int64 `json:"transcode,omitempty"`
}

func (o VodSampleData) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VodSampleData struct{}"
	}

	return strings.Join([]string{"VodSampleData", string(data)}, " ")
}
