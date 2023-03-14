package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageVariationTaskConfigCommon struct {
	Inference *ImageVariationInference `json:"inference"`
}

func (o ImageVariationTaskConfigCommon) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageVariationTaskConfigCommon struct{}"
	}

	return strings.Join([]string{"ImageVariationTaskConfigCommon", string(data)}, " ")
}
