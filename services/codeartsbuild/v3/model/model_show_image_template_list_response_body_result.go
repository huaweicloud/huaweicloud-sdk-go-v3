package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageTemplateListResponseBodyResult 返回结果
type ShowImageTemplateListResponseBodyResult struct {

	// 镜像模版列表
	ImageTemplates *[]ShowImageTemplateListResponseBodyResultImageTemplates `json:"image_templates,omitempty"`
}

func (o ShowImageTemplateListResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageTemplateListResponseBodyResult struct{}"
	}

	return strings.Join([]string{"ShowImageTemplateListResponseBodyResult", string(data)}, " ")
}
