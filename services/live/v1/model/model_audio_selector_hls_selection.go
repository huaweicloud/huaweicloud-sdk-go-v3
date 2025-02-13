package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AudioSelectorHlsSelection hls音频选择器
type AudioSelectorHlsSelection struct {

	// hls音频选择器名
	Name string `json:"name"`

	// hls音频选择器gid
	GroupId string `json:"group_id"`
}

func (o AudioSelectorHlsSelection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioSelectorHlsSelection struct{}"
	}

	return strings.Join([]string{"AudioSelectorHlsSelection", string(data)}, " ")
}
