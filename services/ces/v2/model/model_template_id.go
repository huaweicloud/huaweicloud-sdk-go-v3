package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateId 告警模板的ID，以at开头，后跟字母、数字，长度最长为64
type TemplateId struct {
}

func (o TemplateId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateId struct{}"
	}

	return strings.Join([]string{"TemplateId", string(data)}, " ")
}
