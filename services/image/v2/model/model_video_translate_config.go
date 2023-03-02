package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VideoTranslateConfig struct {
	Common *VideoTranslateConfigCommon `json:"common,omitempty"`
}

func (o VideoTranslateConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoTranslateConfig struct{}"
	}

	return strings.Join([]string{"VideoTranslateConfig", string(data)}, " ")
}
