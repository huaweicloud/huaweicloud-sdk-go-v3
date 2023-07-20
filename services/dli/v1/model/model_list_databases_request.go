package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabasesRequest Request Object
type ListDatabasesRequest struct {

	// 过滤关键字
	Keyword *string `json:"keyword,omitempty"`

	// 查询数量
	Limit *int32 `json:"limit,omitempty"`

	// 查询偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 标签过滤
	Tags *string `json:"tags,omitempty"`

	// 是否返回隐私信息
	WithPriv *bool `json:"with-priv,omitempty"`
}

func (o ListDatabasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabasesRequest struct{}"
	}

	return strings.Join([]string{"ListDatabasesRequest", string(data)}, " ")
}
