package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfirmIaConfigsRequestBody struct {
	// 确认配置项列表

	Configs *[]ConfirmIaConfigRequestBody `json:"configs,omitempty"`
}

func (o ConfirmIaConfigsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmIaConfigsRequestBody struct{}"
	}

	return strings.Join([]string{"ConfirmIaConfigsRequestBody", string(data)}, " ")
}
