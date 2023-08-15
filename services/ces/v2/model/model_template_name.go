package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateName 告警模板的名称，以字母或汉字开头，可包含字母、数字、汉字、_、-，长度范围[1,128]
type TemplateName struct {
}

func (o TemplateName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateName struct{}"
	}

	return strings.Join([]string{"TemplateName", string(data)}, " ")
}
