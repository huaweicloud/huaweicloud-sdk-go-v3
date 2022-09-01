package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchImportConfigRequestBody struct {

	// 配置项ID
	Id string `json:"id" xml:"id"`

	// 配置项名称
	Name string `json:"name" xml:"name"`

	// 配置项详情，长度2MB以内
	Value string `json:"value" xml:"value"`

	// 配置项描述
	Description *string `json:"description,omitempty" xml:"description"`
}

func (o BatchImportConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchImportConfigRequestBody struct{}"
	}

	return strings.Join([]string{"BatchImportConfigRequestBody", string(data)}, " ")
}
