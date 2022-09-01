package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 应用信息
type AppInfo struct {

	// 环境名称
	EnvName *string `json:"env_name,omitempty" xml:"env_name"`

	// 环境id
	EnvId *int64 `json:"env_id,omitempty" xml:"env_id"`

	// 应用空名称
	AppName *string `json:"app_name,omitempty" xml:"app_name"`

	// 应用id
	AppId *int64 `json:"app_id,omitempty" xml:"app_id"`

	// 在线探针数
	OnlineCount *int32 `json:"online_count,omitempty" xml:"online_count"`

	// 手动停止探针数
	DisableCount *int32 `json:"disable_count,omitempty" xml:"disable_count"`

	// 离线探针数
	OfflineCount *int32 `json:"offline_count,omitempty" xml:"offline_count"`
}

func (o AppInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppInfo struct{}"
	}

	return strings.Join([]string{"AppInfo", string(data)}, " ")
}
