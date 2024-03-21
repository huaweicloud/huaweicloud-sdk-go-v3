package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFunctionTemplateRequest Request Object
type ListFunctionTemplateRequest struct {

	// 本次查询起始位置，默认值0
	Marker *string `json:"marker,omitempty"`

	// 每次查询获取的最大模板数量。
	Maxitems *string `json:"maxitems,omitempty"`

	// 是否为公开模板
	Ispublic *string `json:"ispublic,omitempty"`

	// 指定运行时模板
	Runtime *string `json:"runtime,omitempty"`

	// 指定场景模板
	Scene *string `json:"scene,omitempty"`

	// 指定云服务模板
	Service *string `json:"service,omitempty"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ListFunctionTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListFunctionTemplateRequest", string(data)}, " ")
}
