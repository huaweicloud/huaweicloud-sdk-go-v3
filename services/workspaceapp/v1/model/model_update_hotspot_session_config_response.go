package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHotspotSessionConfigResponse Response Object
type UpdateHotspotSessionConfigResponse struct {

	// 唯一标识。
	Id *string `json:"id,omitempty"`

	// 是否开启热点会话迁移。取值为： false：表示关闭。 true：表示开启。
	HotspotSessionMigrationEnable *bool `json:"hotspot_session_migration_enable,omitempty"`

	// 热点时退出会话个数。
	HotspotExitSessionNum *int32 `json:"hotspot_exit_session_num,omitempty"`
	HttpStatusCode        int    `json:"-"`
}

func (o UpdateHotspotSessionConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHotspotSessionConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateHotspotSessionConfigResponse", string(data)}, " ")
}
