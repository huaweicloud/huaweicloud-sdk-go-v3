package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpenEntityHeader struct {

	// 属性
	Attributes *interface{} `json:"attributes,omitempty"`

	// 数据连接
	Connection *[]Connection `json:"connection,omitempty"`

	// 展示文档
	DisplayText *string `json:"display_text,omitempty"`

	// 资产guid
	Guid *string `json:"guid,omitempty"`

	// 类型名称
	TypeName *string `json:"type_name,omitempty"`

	// 标签列表
	Tags *[]TagHeader `json:"tags,omitempty"`

	// 分类名称列表
	ClassificationNames *[]string `json:"classification_names,omitempty"`
}

func (o OpenEntityHeader) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenEntityHeader struct{}"
	}

	return strings.Join([]string{"OpenEntityHeader", string(data)}, " ")
}
