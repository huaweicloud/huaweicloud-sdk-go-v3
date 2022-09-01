package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TransTemplate struct {

	// 转码模板名称。
	TemplateName string `json:"template_name" xml:"template_name"`

	Video *Video `json:"video,omitempty" xml:"video"`

	Audio *Audio `json:"audio,omitempty" xml:"audio"`

	Common *Common `json:"common" xml:"common"`
}

func (o TransTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransTemplate struct{}"
	}

	return strings.Join([]string{"TransTemplate", string(data)}, " ")
}
