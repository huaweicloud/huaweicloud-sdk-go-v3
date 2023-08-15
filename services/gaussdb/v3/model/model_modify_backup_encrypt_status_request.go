package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyBackupEncryptStatusRequest Request Object
type ModifyBackupEncryptStatusRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 租户在某一project下的实例ID。
	InstanceId string `json:"instance_id"`

	Body *BackupEncryptRequest `json:"body,omitempty"`
}

func (o ModifyBackupEncryptStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyBackupEncryptStatusRequest struct{}"
	}

	return strings.Join([]string{"ModifyBackupEncryptStatusRequest", string(data)}, " ")
}
