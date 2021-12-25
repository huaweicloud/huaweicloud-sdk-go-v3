package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateProductTemplateRequestBody struct {
	// 产品模板名称，支持中文,英文大小写，数字，下划线和中划线,长度2-64

	Name string `json:"name"`
	// 产品模板描述，长度0-200

	Description *string `json:"description,omitempty"`
	// 产品模板状态 0-启用 1-禁用

	Status int32 `json:"status"`
}

func (o CreateProductTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProductTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateProductTemplateRequestBody", string(data)}, " ")
}
