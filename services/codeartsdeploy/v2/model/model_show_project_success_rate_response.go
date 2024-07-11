package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectSuccessRateResponse Response Object
type ShowProjectSuccessRateResponse struct {

	// 成功率
	SuccessRate *string `json:"success_rate,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 项目名称
	ProjectName *string `json:"project_name,omitempty"`

	// 应用部署开始时间范围的左边界（包含），格式yyyy-MM-dd
	StartDate *string `json:"start_date,omitempty"`

	// 应用部署开始时间范围的右边界（包含），格式yyyy-MM-dd
	EndDate *string `json:"end_date,omitempty"`

	// 查询到的应用数
	TaskCount *int32 `json:"task_count,omitempty"`

	// 查询到的应用部署记录数
	RecordCount *int32 `json:"record_count,omitempty"`

	// 成功的应用部署记录数
	SuccessRecordCount *int32 `json:"success_record_count,omitempty"`
	HttpStatusCode     int    `json:"-"`
}

func (o ShowProjectSuccessRateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectSuccessRateResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectSuccessRateResponse", string(data)}, " ")
}
