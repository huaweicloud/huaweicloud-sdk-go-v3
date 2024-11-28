package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopBackupRequestBody 停止备份
type StopBackupRequestBody struct {

	// 操作。stop代表停止。
	Action string `json:"action"`
}

func (o StopBackupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopBackupRequestBody struct{}"
	}

	return strings.Join([]string{"StopBackupRequestBody", string(data)}, " ")
}
