package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAtlasesRequest Request Object
type ListAtlasesRequest struct {

	// 分页偏移量，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 每页条数，1~100，默认10
	Limit *int32 `json:"limit,omitempty"`

	// 图库名称（模糊匹配）
	Name *string `json:"name,omitempty"`

	// 图库类型 0=白库 / 1=黑库
	Type *int32 `json:"type,omitempty"`
}

func (o ListAtlasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAtlasesRequest struct{}"
	}

	return strings.Join([]string{"ListAtlasesRequest", string(data)}, " ")
}
