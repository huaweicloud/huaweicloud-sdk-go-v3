package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteBackupErrorResponse struct {

	// 错误码。
	ErrorCode string `json:"error_code"`

	// 错误消息。
	ErrorMsg string `json:"error_msg"`
}

func (o DeleteBackupErrorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackupErrorResponse struct{}"
	}

	return strings.Join([]string{"DeleteBackupErrorResponse", string(data)}, " ")
}
