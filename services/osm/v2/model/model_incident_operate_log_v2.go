package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncidentOperateLogV2 struct {

	// 工单日志id
	IncidentLogId *string `json:"incident_log_id,omitempty"`

	// 工单id
	IncidentId *string `json:"incident_id,omitempty"`

	// 操作类型
	OperateType *int32 `json:"operate_type,omitempty"`

	// 操作员类型
	OperatorType *int32 `json:"operator_type,omitempty"`

	// 操作员id
	OperatorId *string `json:"operator_id,omitempty"`

	// 操作员名称
	OperatorName *string `json:"operator_name,omitempty"`

	// 操作描述
	OperateDesc *string `json:"operate_desc,omitempty"`

	// 操作时间
	OperateTime *string `json:"operate_time,omitempty"`

	// 工单操作时间
	TimestampOperateTime *sdktime.SdkTime `json:"timestamp_operate_time,omitempty"`
}

func (o IncidentOperateLogV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentOperateLogV2 struct{}"
	}

	return strings.Join([]string{"IncidentOperateLogV2", string(data)}, " ")
}
