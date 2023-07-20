package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDliAgencyResponse Response Object
type ShowDliAgencyResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 版本号
	Version *string `json:"version,omitempty"`

	// 当前已有委托
	CurrentRoles   *[]string `json:"current_roles,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowDliAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDliAgencyResponse struct{}"
	}

	return strings.Join([]string{"ShowDliAgencyResponse", string(data)}, " ")
}
