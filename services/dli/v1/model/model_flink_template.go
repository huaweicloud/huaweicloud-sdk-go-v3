package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkTemplate 创建模板信息。
type FlinkTemplate struct {

	// 模板ID
	TemplateId *int64 `json:"template_id,omitempty"`

	// 模板名称。
	Name *string `json:"name,omitempty"`

	// 模板描述。
	Desc *string `json:"desc,omitempty"`

	// 模板创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 作业模板类型。
	JobType *string `json:"job_type,omitempty"`
}

func (o FlinkTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkTemplate struct{}"
	}

	return strings.Join([]string{"FlinkTemplate", string(data)}, " ")
}
