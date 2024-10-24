package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGroupMembershipReqBody 创建成员和组之间的关系。必须指定以下标识符：GroupId、IdentityStoreId和memberId
type CreateGroupMembershipReqBody struct {

	// 身份源中IAM身份中心用户组的全局唯一标识符（ID）
	GroupId string `json:"group_id"`

	MemberId *MemberIdDto `json:"member_id"`
}

func (o CreateGroupMembershipReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupMembershipReqBody struct{}"
	}

	return strings.Join([]string{"CreateGroupMembershipReqBody", string(data)}, " ")
}
