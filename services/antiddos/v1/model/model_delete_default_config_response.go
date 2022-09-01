package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteDefaultConfigResponse struct {

	// 内部错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 内部错误描述
	ErrorMsg       *string `json:"error_msg,omitempty" xml:"error_msg"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDefaultConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDefaultConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteDefaultConfigResponse", string(data)}, " ")
}
