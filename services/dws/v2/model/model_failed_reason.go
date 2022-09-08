package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type FailedReason struct {

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o FailedReason) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailedReason struct{}"
	}

	return strings.Join([]string{"FailedReason", string(data)}, " ")
}
