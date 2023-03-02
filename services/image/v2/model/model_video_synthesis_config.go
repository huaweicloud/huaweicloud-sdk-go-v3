package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VideoSynthesisConfig struct {
	Common *VideoSynthesisConfigCommon `json:"common,omitempty"`
}

func (o VideoSynthesisConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoSynthesisConfig struct{}"
	}

	return strings.Join([]string{"VideoSynthesisConfig", string(data)}, " ")
}
