package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGroupResponse Response Object
type CreateGroupResponse struct {

	// 身份源中IAM身份中心用户组的全局唯一标识符（ID）
	GroupId *string `json:"group_id,omitempty"`

	// 身份源的全局唯一标识符（ID）
	IdentityStoreId *string `json:"identity_store_id,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o CreateGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateGroupResponse", string(data)}, " ")
}
