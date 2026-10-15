package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceHealthReportConfigDto struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 关联的模板ID，未配置时为null。
	TemplateId *int64 `json:"template_id,omitempty"`

	// 关联的模板名称，未配置时为null。
	TemplateName *string `json:"template_name,omitempty"`

	// 关联的模板是否为系统预设模板，未配置时为null。
	System *bool `json:"system,omitempty"`
}

func (o InstanceHealthReportConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceHealthReportConfigDto struct{}"
	}

	return strings.Join([]string{"InstanceHealthReportConfigDto", string(data)}, " ")
}
