package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateServerConfig 节点迁移场景服务器配置
type MigrateServerConfig struct {
	RootVolume *MigrateVolumeSpec `json:"rootVolume,omitempty"`
}

func (o MigrateServerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateServerConfig struct{}"
	}

	return strings.Join([]string{"MigrateServerConfig", string(data)}, " ")
}
