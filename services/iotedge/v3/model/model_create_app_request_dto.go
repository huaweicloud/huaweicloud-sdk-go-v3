package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAppRequestDto struct {

	// 应用ID
	AppId string `json:"app_id"`

	// 应用名称
	AppName *string `json:"app_name,omitempty"`

	// 应用描述
	Description *string `json:"description,omitempty"`

	// 应用类型
	AppType string `json:"app_type"`
}

func (o CreateAppRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppRequestDto struct{}"
	}

	return strings.Join([]string{"CreateAppRequestDto", string(data)}, " ")
}
