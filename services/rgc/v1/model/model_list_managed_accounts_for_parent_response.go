package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListManagedAccountsForParentResponse Response Object
type ListManagedAccountsForParentResponse struct {

	// 纳管的账号信息。
	ManagedAccounts *[]ManagedAccount `json:"managed_accounts,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListManagedAccountsForParentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListManagedAccountsForParentResponse struct{}"
	}

	return strings.Join([]string{"ListManagedAccountsForParentResponse", string(data)}, " ")
}
