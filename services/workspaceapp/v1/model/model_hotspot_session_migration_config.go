package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HotspotSessionMigrationConfig struct {

	// 唯一标识。
	Id *string `json:"id,omitempty"`

	// 项目id。
	ProjectId *string `json:"project_id,omitempty"`

	// 是否开启热点会话迁移。取值为： false：表示关闭。 true：表示开启。
	HotspotSessionMigrationEnable *bool `json:"hotspot_session_migration_enable,omitempty"`

	// 热点时退出会话个数。
	HotspotExitSessionNum *int32 `json:"hotspot_exit_session_num,omitempty"`

	// 热点时不迁移用户id列表。
	NonMigrateUsers *[]UserInfo `json:"non_migrate_users,omitempty"`
}

func (o HotspotSessionMigrationConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HotspotSessionMigrationConfig struct{}"
	}

	return strings.Join([]string{"HotspotSessionMigrationConfig", string(data)}, " ")
}
