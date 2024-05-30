package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EngineCreateReqMaintenanceConfig 微服务引擎的维护时间窗
type EngineCreateReqMaintenanceConfig struct {

	// 维护时间
	Time *string `json:"time,omitempty"`

	// 维护时间的时区
	Zone *string `json:"zone,omitempty"`
}

func (o EngineCreateReqMaintenanceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineCreateReqMaintenanceConfig struct{}"
	}

	return strings.Join([]string{"EngineCreateReqMaintenanceConfig", string(data)}, " ")
}
