package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageVariationTaskConfig struct {
	Common *ImageVariationTaskConfigCommon `json:"common"`
}

func (o ImageVariationTaskConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageVariationTaskConfig struct{}"
	}

	return strings.Join([]string{"ImageVariationTaskConfig", string(data)}, " ")
}
