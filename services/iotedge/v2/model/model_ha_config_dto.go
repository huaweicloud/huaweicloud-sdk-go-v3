package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HaConfigDto 高可用配置请求结构体
type HaConfigDto struct {

	// 节点高可用类型双活或者主备
	HaType *string `json:"ha_type,omitempty"`

	ActiveStandbyConfig *ActiveStandbyConfigDto `json:"active_standby_config,omitempty"`
}

func (o HaConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HaConfigDto struct{}"
	}

	return strings.Join([]string{"HaConfigDto", string(data)}, " ")
}
