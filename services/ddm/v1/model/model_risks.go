package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Risks struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 引擎名称。
	EngineName string `json:"engine_name"`

	// 当前引擎版本。
	EngineVersion string `json:"engine_version"`

	// 风险等级。
	Level *int32 `json:"level,omitempty"`

	// 建议升级原因。
	Suggest *string `json:"suggest,omitempty"`

	// 升级影响。
	Influence *string `json:"influence,omitempty"`

	// 指导连接。
	Guidance *string `json:"guidance,omitempty"`

	// 业务影响时长。
	ServiceImpactDuration *string `json:"service_impact_duration,omitempty"`

	// 升级时长。
	UpgradeDuration *string `json:"upgrade_duration,omitempty"`
}

func (o Risks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Risks struct{}"
	}

	return strings.Join([]string{"Risks", string(data)}, " ")
}
