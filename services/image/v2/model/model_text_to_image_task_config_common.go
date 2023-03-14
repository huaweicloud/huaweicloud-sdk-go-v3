package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TextToImageTaskConfigCommon struct {
	Inference *TextToImageInference `json:"inference"`
}

func (o TextToImageTaskConfigCommon) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextToImageTaskConfigCommon struct{}"
	}

	return strings.Join([]string{"TextToImageTaskConfigCommon", string(data)}, " ")
}
