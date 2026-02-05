package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHotspotSessionConfigReq 创建热点会话迁移配置请求。
type CreateHotspotSessionConfigReq struct {

	// 是否开启热点会话迁移。取值为： false：表示关闭。 true：表示开启。
	HotspotSessionMigrationEnable bool `json:"hotspot_session_migration_enable"`

	// 热点时退出会话个数。
	HotspotExitSessionNum int32 `json:"hotspot_exit_session_num"`

	// 热点时不迁移用户id列表。
	NonMigrateUsers *[]UserInfo `json:"non_migrate_users,omitempty"`
}

func (o CreateHotspotSessionConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHotspotSessionConfigReq struct{}"
	}

	return strings.Join([]string{"CreateHotspotSessionConfigReq", string(data)}, " ")
}
