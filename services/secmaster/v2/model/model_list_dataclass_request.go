package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataclassRequest Request Object
type ListDataclassRequest struct {

	// 内容类型
	ContentType string `json:"content-type"`

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 数据量
	Limit *int32 `json:"limit,omitempty"`

	// 名称查询
	Name *string `json:"name,omitempty"`

	// 业务编码
	BusinessCode *string `json:"business_code,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 是否内置
	IsBuiltIn *bool `json:"is_built_in,omitempty"`
}

func (o ListDataclassRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataclassRequest struct{}"
	}

	return strings.Join([]string{"ListDataclassRequest", string(data)}, " ")
}
