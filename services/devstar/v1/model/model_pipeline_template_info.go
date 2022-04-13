package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineTemplateInfo struct {
	// 流水线模板的id。

	Id *string `json:"id,omitempty"`
	// 流水线模板的名称。

	Name *string `json:"name,omitempty"`
	// 流水线模板的详细信息。

	Detail *string `json:"detail,omitempty"`
}

func (o PipelineTemplateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineTemplateInfo struct{}"
	}

	return strings.Join([]string{"PipelineTemplateInfo", string(data)}, " ")
}
