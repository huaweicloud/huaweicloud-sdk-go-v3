package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateGroupUserGroupResponse Response Object
type AssociateGroupUserGroupResponse struct {

	// 关联成功成员列表
	Success *[]MemberBaseDto `json:"success,omitempty"`

	// 关联失败成员列表
	Failure        *[]FailureForBatchCreateGroupMembersDto `json:"failure,omitempty"`
	HttpStatusCode int                                     `json:"-"`
}

func (o AssociateGroupUserGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateGroupUserGroupResponse struct{}"
	}

	return strings.Join([]string{"AssociateGroupUserGroupResponse", string(data)}, " ")
}
