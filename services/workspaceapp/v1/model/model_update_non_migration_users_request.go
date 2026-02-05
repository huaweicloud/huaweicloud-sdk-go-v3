package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNonMigrationUsersRequest Request Object
type UpdateNonMigrationUsersRequest struct {

	// 热点会话迁移配置ID。
	ConfigId string `json:"config_id"`

	Body *UpdateNonMigrationUsersReq `json:"body,omitempty"`
}

func (o UpdateNonMigrationUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNonMigrationUsersRequest struct{}"
	}

	return strings.Join([]string{"UpdateNonMigrationUsersRequest", string(data)}, " ")
}
