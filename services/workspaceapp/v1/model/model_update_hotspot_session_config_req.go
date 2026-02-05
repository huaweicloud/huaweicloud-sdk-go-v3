package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHotspotSessionConfigReq 更新热点会话迁移配置请求体。
type UpdateHotspotSessionConfigReq struct {

	// 是否开启热点会话迁移。取值为： false：表示关闭。 true：表示开启。
	HotspotSessionMigrationEnable bool `json:"hotspot_session_migration_enable"`

	// 热点时退出会话个数。
	HotspotExitSessionNum int32 `json:"hotspot_exit_session_num"`
}

func (o UpdateHotspotSessionConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHotspotSessionConfigReq struct{}"
	}

	return strings.Join([]string{"UpdateHotspotSessionConfigReq", string(data)}, " ")
}
