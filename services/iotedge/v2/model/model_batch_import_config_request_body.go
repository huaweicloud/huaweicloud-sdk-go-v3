package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchImportConfigRequestBody struct {

	// 配置项ID
	Id string `json:"id"`

	// 配置项名称
	Name string `json:"name"`

	// 配置项详情，长度2MB以内
	Value string `json:"value"`

	// 配置项描述
	Description *string `json:"description,omitempty"`
}

func (o BatchImportConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchImportConfigRequestBody struct{}"
	}

	return strings.Join([]string{"BatchImportConfigRequestBody", string(data)}, " ")
}
