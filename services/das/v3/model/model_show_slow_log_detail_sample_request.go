package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowLogDetailSampleRequest Request Object
type ShowSlowLogDetailSampleRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 开始时间（Unix timestamp），单位：毫秒
	StartTime int64 `json:"start_time"`

	// 结束时间（Unix timestamp），单位：毫秒
	EndTime int64 `json:"end_time"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`

	// SQL模板ID
	SqlTemplateId string `json:"sql_template_id"`

	// 是否需要数据库名
	WithDb *string `json:"with_db,omitempty"`
}

func (o ShowSlowLogDetailSampleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowLogDetailSampleRequest struct{}"
	}

	return strings.Join([]string{"ShowSlowLogDetailSampleRequest", string(data)}, " ")
}
