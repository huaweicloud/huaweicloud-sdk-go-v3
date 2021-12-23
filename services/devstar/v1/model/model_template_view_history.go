package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateViewHistory struct {
	// 模板的id。

	TemplateId string `json:"template_id"`
	// 模板的名称。

	TemplateTitle string `json:"template_title"`
}

func (o TemplateViewHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateViewHistory struct{}"
	}

	return strings.Join([]string{"TemplateViewHistory", string(data)}, " ")
}
