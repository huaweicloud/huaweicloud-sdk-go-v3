package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDatabaseLevelDatabaseRequest Request Object
type ShowDatabaseLevelDatabaseRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 备份id
	BackupId string `json:"backup_id"`

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowDatabaseLevelDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatabaseLevelDatabaseRequest struct{}"
	}

	return strings.Join([]string{"ShowDatabaseLevelDatabaseRequest", string(data)}, " ")
}
