package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateLabelsResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 错误描述
	ErrorMsg       *string `json:"error_msg,omitempty" xml:"error_msg"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLabelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLabelsResponse struct{}"
	}

	return strings.Join([]string{"CreateLabelsResponse", string(data)}, " ")
}
