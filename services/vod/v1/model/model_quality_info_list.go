package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QualityInfoList struct {
	Video *VideoInfo `json:"video"`

	Audio *AudioInfo `json:"audio,omitempty"`
}

func (o QualityInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QualityInfoList struct{}"
	}

	return strings.Join([]string{"QualityInfoList", string(data)}, " ")
}
