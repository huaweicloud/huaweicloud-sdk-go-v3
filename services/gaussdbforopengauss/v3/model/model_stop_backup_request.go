package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopBackupRequest Request Object
type StopBackupRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o StopBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopBackupRequest struct{}"
	}

	return strings.Join([]string{"StopBackupRequest", string(data)}, " ")
}
