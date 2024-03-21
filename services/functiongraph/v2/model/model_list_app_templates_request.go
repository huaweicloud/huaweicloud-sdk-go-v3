package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppTemplatesRequest Request Object
type ListAppTemplatesRequest struct {

	// 本次查询起始位置，默认值0
	Marker *string `json:"marker,omitempty"`

	// 本次查询最大返回的数据条数，最大值500，默认值100
	Maxitems *string `json:"maxitems,omitempty"`

	// 模板执行运行时
	Runtime *string `json:"runtime,omitempty"`

	// 模板类别
	Category *string `json:"category,omitempty"`

	// 模板语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ListAppTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListAppTemplatesRequest", string(data)}, " ")
}
