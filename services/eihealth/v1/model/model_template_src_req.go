package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateSrcReq struct {

	// 源模板id
	SourceTemplateId string `json:"source_template_id"`

	// 导入模板名称
	DestinationTemplateName string `json:"destination_template_name"`
}

func (o TemplateSrcReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateSrcReq struct{}"
	}

	return strings.Join([]string{"TemplateSrcReq", string(data)}, " ")
}
