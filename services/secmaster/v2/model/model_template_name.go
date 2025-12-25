package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateName 模板名称
type TemplateName struct {
}

func (o TemplateName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateName struct{}"
	}

	return strings.Join([]string{"TemplateName", string(data)}, " ")
}
