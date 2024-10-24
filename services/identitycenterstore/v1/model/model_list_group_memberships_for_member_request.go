package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupMembershipsForMemberRequest Request Object
type ListGroupMembershipsForMemberRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 身份源的全局唯一标识符（ID）
	IdentityStoreId string `json:"identity_store_id"`

	// 每个请求返回的最大结果数
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记
	Marker *string `json:"marker,omitempty"`

	// 用户唯一标识
	UserId string `json:"user_id"`
}

func (o ListGroupMembershipsForMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupMembershipsForMemberRequest struct{}"
	}

	return strings.Join([]string{"ListGroupMembershipsForMemberRequest", string(data)}, " ")
}
