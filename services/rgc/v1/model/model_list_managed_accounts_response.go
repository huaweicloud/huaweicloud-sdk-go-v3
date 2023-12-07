package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListManagedAccountsResponse Response Object
type ListManagedAccountsResponse struct {

	// 纳管的Account。
	ManagedAccounts *[]ManagedAccount `json:"managed_accounts,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListManagedAccountsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListManagedAccountsResponse struct{}"
	}

	return strings.Join([]string{"ListManagedAccountsResponse", string(data)}, " ")
}
