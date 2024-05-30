package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectErrorInfo 错误信息列表。
type ConnectErrorInfo struct {

	// 云手机的唯一标识ID。
	PhoneId *string `json:"phone_id,omitempty"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误说明。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o ConnectErrorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectErrorInfo struct{}"
	}

	return strings.Join([]string{"ConnectErrorInfo", string(data)}, " ")
}
