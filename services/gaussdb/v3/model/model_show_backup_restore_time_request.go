package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupRestoreTimeRequest Request Object
type ShowBackupRestoreTimeRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 租户在某一project下的实例ID。
	InstanceId string `json:"instance_id"`

	// 所需查询的日志，为yyyy-mm-dd字符串格式，时区为UTC。
	Date *string `json:"date,omitempty"`
}

func (o ShowBackupRestoreTimeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupRestoreTimeRequest struct{}"
	}

	return strings.Join([]string{"ShowBackupRestoreTimeRequest", string(data)}, " ")
}
