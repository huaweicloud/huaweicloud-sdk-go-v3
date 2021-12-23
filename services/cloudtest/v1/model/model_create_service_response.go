package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateServiceResponse struct {
	// 接口调用成功返回的服务名

	ServiceName *string `json:"service_name,omitempty"`
	// 接口调用成功返回的服务id

	ServiceId *int32 `json:"service_id,omitempty"`
	// 接口调用成功不返回，调用失败错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 接口调用成功不返回，调用失败错误信息

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServiceResponse struct{}"
	}

	return strings.Join([]string{"CreateServiceResponse", string(data)}, " ")
}
