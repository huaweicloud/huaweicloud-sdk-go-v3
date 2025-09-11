package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateReportRequestBody struct {

	// 数据库ID列表，多个用英文逗号分割
	DbIds *string `json:"db_ids,omitempty"`

	// 结束日期
	EndDate *sdktime.SdkTime `json:"end_date,omitempty"`

	// 结束时间
	EndTime string `json:"end_time"`

	// 开始日期
	StartDate *sdktime.SdkTime `json:"start_date,omitempty"`

	// 开始时间
	StartTime string `json:"start_time"`

	// 模板ID
	TemplateId string `json:"template_id"`
}

func (o CreateReportRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReportRequestBody struct{}"
	}

	return strings.Join([]string{"CreateReportRequestBody", string(data)}, " ")
}
