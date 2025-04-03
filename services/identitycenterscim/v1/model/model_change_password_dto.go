package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangePasswordDto struct {

	// 一个布尔值，表示服务提供商是否支持这种操作
	Supported *bool `json:"supported,omitempty"`
}

func (o ChangePasswordDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePasswordDto struct{}"
	}

	return strings.Join([]string{"ChangePasswordDto", string(data)}, " ")
}
