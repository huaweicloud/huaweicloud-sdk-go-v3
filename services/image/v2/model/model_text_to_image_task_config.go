package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TextToImageTaskConfig struct {
	Common *TextToImageTaskConfigCommon `json:"common"`
}

func (o TextToImageTaskConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextToImageTaskConfig struct{}"
	}

	return strings.Join([]string{"TextToImageTaskConfig", string(data)}, " ")
}
