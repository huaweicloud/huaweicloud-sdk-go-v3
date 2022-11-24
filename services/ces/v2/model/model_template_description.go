package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 告警模板的描述，长度范围[0,256]，该字段默认值为空字符串
type TemplateDescription struct {
}

func (o TemplateDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateDescription struct{}"
	}

	return strings.Join([]string{"TemplateDescription", string(data)}, " ")
}
