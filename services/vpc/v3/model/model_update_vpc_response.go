package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateVpcResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	// 错误消息
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`

	// 错误码
	ErrorCode      *string `json:"error_code,omitempty" xml:"error_code"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateVpcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcResponse struct{}"
	}

	return strings.Join([]string{"UpdateVpcResponse", string(data)}, " ")
}
