package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnhancedConnectionsRequest Request Object
type ListEnhancedConnectionsRequest struct {

	// 查询最大连接个数，默认100。
	Limit *int32 `json:"limit,omitempty"`

	// 连接名。
	Name *string `json:"name,omitempty"`

	// 查询结果偏移量，默认为0（连接以创建时间进行排序）
	Offset *int32 `json:"offset,omitempty"`

	// 连接状态，包括以下两种状态： ACTIVE：已激活 DELETED：已删除
	Status *string `json:"status,omitempty"`

	// 标签
	Tags *string `json:"tags,omitempty"`
}

func (o ListEnhancedConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnhancedConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListEnhancedConnectionsRequest", string(data)}, " ")
}
