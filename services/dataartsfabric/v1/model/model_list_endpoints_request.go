package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEndpointsRequest Request Object
type ListEndpointsRequest struct {

	// Workspace的ID
	WorkspaceId string `json:"workspace_id"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 指定每一页返回的最大条目数，取值范围[1,100]，默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// 通过模型ID检索，32~36位的英文、数字、短横组合
	ModelId *string `json:"model_id,omitempty"`

	// 通过名字搜索Endpoint的参数，一个Endpoint名称，只能包含中文、字母、数字、下划线、中划线，支持模糊查询
	Name *string `json:"name,omitempty"`

	// 通过id检索Endpoint的参数
	EndpointId *string `json:"endpoint_id,omitempty"`

	// 通过类型检索Endpoint的参数
	Type *EndpointType `json:"type,omitempty"`

	// 可见性检索的参数，可选值为： - PRIVATE: 私有，用户自己创建的； - PUBLIC:公共，查询所有公共的，包括其他用户创建的； - ALL: 所有的； - 默认为空，不填表示不限制，则查出当前用户下的，包括PRIVATE和PUBLIC，不包括其他用户创建的。
	Visibility *string `json:"visibility,omitempty"`
}

func (o ListEndpointsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointsRequest struct{}"
	}

	return strings.Join([]string{"ListEndpointsRequest", string(data)}, " ")
}
