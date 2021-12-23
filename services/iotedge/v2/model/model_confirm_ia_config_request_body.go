package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfirmIaConfigRequestBody struct {
	// 配置项ID

	Id string `json:"id"`
	// 版本号

	Version string `json:"version"`
}

func (o ConfirmIaConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmIaConfigRequestBody struct{}"
	}

	return strings.Join([]string{"ConfirmIaConfigRequestBody", string(data)}, " ")
}
