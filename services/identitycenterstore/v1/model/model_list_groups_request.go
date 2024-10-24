package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupsRequest Request Object
type ListGroupsRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 身份源的全局唯一标识符（ID）
	IdentityStoreId string `json:"identity_store_id"`

	// 分页标记
	Marker *string `json:"marker,omitempty"`

	// 每个请求返回的最大结果数
	Limit *int32 `json:"limit,omitempty"`

	// 通过显示名模糊查询用户组信息
	DisplayName *string `json:"display_name,omitempty"`
}

func (o ListGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListGroupsRequest", string(data)}, " ")
}
