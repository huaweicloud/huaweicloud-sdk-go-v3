package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlinkSqlJobTemplate struct {

	// 模板ID。
	TemplateId *int32 `json:"template_id,omitempty"`

	// 模板名称。
	Name *string `json:"name,omitempty"`

	// 模板描述。
	Desc *string `json:"desc,omitempty"`

	// 模板创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 模板更新时间。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// Stream SQL语句, 至少包含source, query, sink三个部分。
	SqlBody *string `json:"sql_body,omitempty"`

	// 作业模板的类型。
	JobType *string `json:"job_type,omitempty"`
}

func (o FlinkSqlJobTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkSqlJobTemplate struct{}"
	}

	return strings.Join([]string{"FlinkSqlJobTemplate", string(data)}, " ")
}
