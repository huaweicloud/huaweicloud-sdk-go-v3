package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageTranslateConfig struct {
	Common *ImageTranslateConfigCommon `json:"common"`
}

func (o ImageTranslateConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageTranslateConfig struct{}"
	}

	return strings.Join([]string{"ImageTranslateConfig", string(data)}, " ")
}
