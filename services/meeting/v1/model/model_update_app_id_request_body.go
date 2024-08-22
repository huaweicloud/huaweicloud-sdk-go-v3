package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppIdRequestBody 修改企业应用的入参
type UpdateAppIdRequestBody struct {

	// 企业应用名称
	AppName *string `json:"app_name,omitempty"`

	// 企业应用描述
	Description *string `json:"description,omitempty"`

	// 企业应用状态  * 0：正常  * 1：停用
	Status *int32 `json:"status,omitempty"`
}

func (o UpdateAppIdRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppIdRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAppIdRequestBody", string(data)}, " ")
}
