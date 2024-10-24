package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListModelsRequest Request Object
type ListModelsRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 指定每一页返回的最大条目数，取值范围[1,100]，默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// 通过名字搜索Model，支持模糊查询
	Name *string `json:"name,omitempty"`

	// 通过模型ID检索，32~36位的英文、数字、短横组合
	Id *string `json:"id,omitempty"`

	// 通过模型类型检索
	Type *string `json:"type,omitempty"`

	// 可见性检索的参数，可选值为： - PRIVATE: 私有，用户自己创建的； - PUBLIC:公共，查询所有公共的，包括其他用户创建的； - ALL: 所有的； - 默认为空，不填表示不限制，则查出当前用户下的，包括PRIVATE和PUBLIC，不包括其他用户创建的。
	Visibility *string `json:"visibility,omitempty"`

	// 指定排序字段，可选值为： CREATE_TIME：创建时间，默认值 UPDATE_TIME：更新时间 NAME: 服务名称
	SortBy *string `json:"sort_by,omitempty"`

	// 排序方式，可选值如下： ASC : 递增排序 DESC: 递减排序，默认值
	OrderBy *string `json:"order_by,omitempty"`
}

func (o ListModelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModelsRequest struct{}"
	}

	return strings.Join([]string{"ListModelsRequest", string(data)}, " ")
}
