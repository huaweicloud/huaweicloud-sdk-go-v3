package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageHighresolutionMattingConfig struct {
	Common *ImageHighresolutionMattingConfigCommon `json:"common"`
}

func (o ImageHighresolutionMattingConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageHighresolutionMattingConfig struct{}"
	}

	return strings.Join([]string{"ImageHighresolutionMattingConfig", string(data)}, " ")
}
