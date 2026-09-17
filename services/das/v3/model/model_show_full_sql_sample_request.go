package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFullSqlSampleRequest Request Object
type ShowFullSqlSampleRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// SQL模板ID
	SqlTemplateId string `json:"sql_template_id"`

	// 开始时间戳（ms）
	StartAt int64 `json:"start_at"`

	// 结束时间戳（ms）
	EndAt int64 `json:"end_at"`
}

func (o ShowFullSqlSampleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFullSqlSampleRequest struct{}"
	}

	return strings.Join([]string{"ShowFullSqlSampleRequest", string(data)}, " ")
}
