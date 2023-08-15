package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageHighresolutionMattingConfigCommon struct {
	Inference *ImageHighresolutionMattingInference `json:"inference"`
}

func (o ImageHighresolutionMattingConfigCommon) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageHighresolutionMattingConfigCommon struct{}"
	}

	return strings.Join([]string{"ImageHighresolutionMattingConfigCommon", string(data)}, " ")
}
