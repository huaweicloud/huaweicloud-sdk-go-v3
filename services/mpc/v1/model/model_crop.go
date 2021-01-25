package model

import (
	"encoding/json"

	"strings"
)

type Crop struct {
	// 截取的视频时长。  单位：秒  从0秒开始算起
	Duration *int32 `json:"duration,omitempty"`
}

func (o Crop) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Crop struct{}"
	}

	return strings.Join([]string{"Crop", string(data)}, " ")
}
