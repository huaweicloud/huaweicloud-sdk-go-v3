package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateJobNameResponse Response Object
type ValidateJobNameResponse struct {

	// 任务名称是否有效。
	IsValid *bool `json:"is_valid,omitempty"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息。
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ValidateJobNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateJobNameResponse struct{}"
	}

	return strings.Join([]string{"ValidateJobNameResponse", string(data)}, " ")
}
