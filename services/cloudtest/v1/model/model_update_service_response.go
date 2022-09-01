package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateServiceResponse struct {

	// 接口调用成功返回的服务名
	ServiceName *string `json:"service_name,omitempty" xml:"service_name"`

	// 接口调用成功返回的服务id
	ServiceId *int32 `json:"service_id,omitempty" xml:"service_id"`

	// 接口调用成功不返回，调用失败错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 接口调用成功不返回，调用失败错误信息
	ErrorMsg       *string `json:"error_msg,omitempty" xml:"error_msg"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServiceResponse struct{}"
	}

	return strings.Join([]string{"UpdateServiceResponse", string(data)}, " ")
}
