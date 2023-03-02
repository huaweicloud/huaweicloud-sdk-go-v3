package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageTranslateConfigCommon struct {
	Inference *ImageTranslateInference `json:"inference"`
}

func (o ImageTranslateConfigCommon) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageTranslateConfigCommon struct{}"
	}

	return strings.Join([]string{"ImageTranslateConfigCommon", string(data)}, " ")
}
