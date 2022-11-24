package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateBodyPrimitiveTypeHolder struct {

	// HCL模板，描述了资源的目标状态。RF将比较此模板与当前远程资源的状态之间的区别。  template_body和template_uri 必须有且只有一个存在
	TemplateBody *string `json:"template_body,omitempty"`
}

func (o TemplateBodyPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateBodyPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"TemplateBodyPrimitiveTypeHolder", string(data)}, " ")
}
