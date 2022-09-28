package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImportTemplateResultRsp struct {

	// 源项目id
	SourceProjectId *string `json:"source_project_id,omitempty"`

	// 源模板id
	SourceTemplateId *string `json:"source_template_id,omitempty"`

	// 导入后的模板id
	DestinationTemplateId *string `json:"destination_template_id,omitempty"`

	// 导入后的模板名称
	DestinationTemplateName *string `json:"destination_template_name,omitempty"`

	// 失败原因，导入失败会返回
	FailedReason *string `json:"failed_reason,omitempty"`

	// 导入结果
	Status *string `json:"status,omitempty"`
}

func (o ImportTemplateResultRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportTemplateResultRsp struct{}"
	}

	return strings.Join([]string{"ImportTemplateResultRsp", string(data)}, " ")
}
