package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowRiskInfoEngineRiskDesc struct {

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 当前引擎
	EngineName *string `json:"engine_name,omitempty"`

	// 当前引擎小版本
	EngineVersion *string `json:"engine_version,omitempty"`

	// 风险等级，默认1
	Level *int64 `json:"level,omitempty"`

	// 建议升级原因,无风险为空
	Suggest *string `json:"suggest,omitempty"`

	// 升级影响，无风险为空
	Influence *string `json:"influence,omitempty"`

	// 指导链接，无风险为空
	Guidance *string `json:"guidance,omitempty"`

	// 业务影响时长，无风险为空
	ServiceImpactDuration *string `json:"service_impact_duration,omitempty"`

	// 升级时长，无风险为空
	UpgradeDuration *string `json:"upgrade_duration,omitempty"`
}

func (o ShowRiskInfoEngineRiskDesc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRiskInfoEngineRiskDesc struct{}"
	}

	return strings.Join([]string{"ShowRiskInfoEngineRiskDesc", string(data)}, " ")
}
