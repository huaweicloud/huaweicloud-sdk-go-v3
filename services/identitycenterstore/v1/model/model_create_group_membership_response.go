package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGroupMembershipResponse Response Object
type CreateGroupMembershipResponse struct {

	// 身份源的全局唯一标识符（ID）
	IdentityStoreId *string `json:"identity_store_id,omitempty"`

	// 身份源中用户和组关联关系的全局唯一标识符（ID）
	MembershipId   *string `json:"membership_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGroupMembershipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupMembershipResponse struct{}"
	}

	return strings.Join([]string{"CreateGroupMembershipResponse", string(data)}, " ")
}
