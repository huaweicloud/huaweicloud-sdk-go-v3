package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribeGroupMembershipRequest Request Object
type DescribeGroupMembershipRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 身份源的全局唯一标识符（ID）
	IdentityStoreId string `json:"identity_store_id"`

	// 身份源中用户和组关联关系的全局唯一标识符（ID）
	MembershipId string `json:"membership_id"`
}

func (o DescribeGroupMembershipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeGroupMembershipRequest struct{}"
	}

	return strings.Join([]string{"DescribeGroupMembershipRequest", string(data)}, " ")
}
