package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHotspotSessionConfigResponse Response Object
type CreateHotspotSessionConfigResponse struct {

	// 唯一标识。
	Id *string `json:"id,omitempty"`

	// 是否开启热点会话迁移。取值为： false：表示关闭。 true：表示开启。
	HotspotSessionMigrationEnable *bool `json:"hotspot_session_migration_enable,omitempty"`

	// 热点时退出会话个数。
	HotspotExitSessionNum *int32 `json:"hotspot_exit_session_num,omitempty"`

	// 热点时不迁移用户id列表。
	NonMigrateUsers *[]UserInfo `json:"non_migrate_users,omitempty"`
	HttpStatusCode  int         `json:"-"`
}

func (o CreateHotspotSessionConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHotspotSessionConfigResponse struct{}"
	}

	return strings.Join([]string{"CreateHotspotSessionConfigResponse", string(data)}, " ")
}
