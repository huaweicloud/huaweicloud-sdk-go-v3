package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLogBackupResponse Response Object
type ShowLogBackupResponse struct {
	LogList *[]LogList `json:"logList,omitempty"`

	// 查询日志的类型。
	Type           *string `json:"type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLogBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogBackupResponse struct{}"
	}

	return strings.Join([]string{"ShowLogBackupResponse", string(data)}, " ")
}
