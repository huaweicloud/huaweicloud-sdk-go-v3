package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAclAccountsResponse Response Object
type ListAclAccountsResponse struct {

	// ACL账号列表
	Accounts       *[]AclAccountResp `json:"accounts,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListAclAccountsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAclAccountsResponse struct{}"
	}

	return strings.Join([]string{"ListAclAccountsResponse", string(data)}, " ")
}
