package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AsyncCreateJobReq 异步创建任务请求体。
type AsyncCreateJobReq struct {
	BaseInfo *JobBaseInfo `json:"base_info"`

	// 创建任务数据库信息体。
	SourceEndpoint []JobEndpointInfo `json:"source_endpoint"`

	// 创建任务数据库信息体。
	TargetEndpoint []JobEndpointInfo `json:"target_endpoint"`

	AlarmNotify *AlarmNotifyConfig `json:"alarm_notify,omitempty"`

	// 限速信息体。 - 限速：自定义的最大迁移速度，迁移过程中的迁移速度将不会超过该速度。 - 不限速：对迁移速度不进行限制，通常会最大化使用源数据库的出口带宽。该流速模式同时会对源数据库造成读消耗，消耗取决于源数据库的出口带宽。比如：源数据库的出口带宽为100MB/s，假设高速模式使用了80%带宽，则迁移对源数据库将造成80MB/s的读操作IO消耗。
	SpeedLimit *[]SpeedLimitInfo `json:"speed_limit,omitempty"`

	UserMigration *UserMigrationInfo `json:"user_migration,omitempty"`

	PolicyConfig *PolicyConfig `json:"policy_config"`

	DbObject *DbObject `json:"db_object"`

	DbParam *DbParamInfo `json:"db_param,omitempty"`

	TuningParams *TuningParamInfo `json:"tuning_params,omitempty"`

	PeriodOrder *PeriodOrderInfo `json:"period_order,omitempty"`

	NodeInfo *JobNodeInfo `json:"node_info"`

	// 指定公网IP的信息
	PublicIpList *[]PublicIpConfig `json:"public_ip_list,omitempty"`
}

func (o AsyncCreateJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AsyncCreateJobReq struct{}"
	}

	return strings.Join([]string{"AsyncCreateJobReq", string(data)}, " ")
}
