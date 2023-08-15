package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryPreCheckResult 预检查结果信息体。
type QueryPreCheckResult struct {

	// 返回的预检查结果是否通过。
	Result *bool `json:"result,omitempty"`

	// 预检查进度百分比。
	Process *string `json:"process,omitempty"`

	// 预检查通过百分比。
	TotalPassedRate *string `json:"total_passed_rate,omitempty"`

	// 数据库实例ID。
	RdsInstanceId *string `json:"rds_instance_id,omitempty"`

	// 迁移方向。
	JobDirection *string `json:"job_direction,omitempty"`

	// 预检查各项结果。
	PrecheckResults *[]PrecheckResult `json:"precheck_results,omitempty"`
}

func (o QueryPreCheckResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryPreCheckResult struct{}"
	}

	return strings.Join([]string{"QueryPreCheckResult", string(data)}, " ")
}
