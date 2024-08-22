package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAppIdRequestBody 添加企业应用入参
type AddAppIdRequestBody struct {

	// 企业应用名称
	AppName string `json:"app_name"`

	// 企业应用描述
	Description *string `json:"description,omitempty"`
}

func (o AddAppIdRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAppIdRequestBody struct{}"
	}

	return strings.Join([]string{"AddAppIdRequestBody", string(data)}, " ")
}
