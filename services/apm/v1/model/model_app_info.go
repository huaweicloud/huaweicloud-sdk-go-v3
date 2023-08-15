package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppInfo 应用信息。
type AppInfo struct {

	// 环境名称。
	EnvName *string `json:"env_name,omitempty"`

	// 环境id。
	EnvId *int64 `json:"env_id,omitempty"`

	// 组件名称。
	AppName *string `json:"app_name,omitempty"`

	// 组件id。
	AppId *int64 `json:"app_id,omitempty"`

	// 在线探针数。
	OnlineCount *int32 `json:"online_count,omitempty"`

	// 手动停止探针数。
	DisableCount *int32 `json:"disable_count,omitempty"`

	// 离线探针数。
	OfflineCount *int32 `json:"offline_count,omitempty"`
}

func (o AppInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppInfo struct{}"
	}

	return strings.Join([]string{"AppInfo", string(data)}, " ")
}
