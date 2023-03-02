package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageToVideoConfig struct {
	Common *ImageToVideoConfigCommon `json:"common,omitempty"`
}

func (o ImageToVideoConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageToVideoConfig struct{}"
	}

	return strings.Join([]string{"ImageToVideoConfig", string(data)}, " ")
}
