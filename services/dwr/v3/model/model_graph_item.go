package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GraphItem struct {

	// 工作流的名称。
	Name *string `json:"name,omitempty"`

	// 系统记录的创建工作流模板的时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 工作流的URN。
	GraphUrn *string `json:"graph_urn,omitempty"`

	// 工作流ID
	Id *string `json:"id,omitempty"`
}

func (o GraphItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GraphItem struct{}"
	}

	return strings.Join([]string{"GraphItem", string(data)}, " ")
}
