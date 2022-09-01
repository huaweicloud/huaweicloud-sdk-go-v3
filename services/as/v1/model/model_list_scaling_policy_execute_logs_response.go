package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListScalingPolicyExecuteLogsResponse struct {

	// 总记录数。
	TotalNumber *int32 `json:"total_number,omitempty" xml:"total_number"`

	// 查询的起始行号。
	StartNumber *int32 `json:"start_number,omitempty" xml:"start_number"`

	// 查询记录数。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 伸缩策略执行日志列表。
	ScalingPolicyExecuteLog *[]ScalingPolicyExecuteLogList `json:"scaling_policy_execute_log,omitempty" xml:"scaling_policy_execute_log"`
	HttpStatusCode          int                            `json:"-"`
}

func (o ListScalingPolicyExecuteLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingPolicyExecuteLogsResponse struct{}"
	}

	return strings.Join([]string{"ListScalingPolicyExecuteLogsResponse", string(data)}, " ")
}
