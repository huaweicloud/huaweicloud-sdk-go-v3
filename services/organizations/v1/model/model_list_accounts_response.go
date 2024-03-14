package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccountsResponse Response Object
type ListAccountsResponse struct {

	// 组织中的账号列表。
	Accounts *[]AccountDto `json:"accounts,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListAccountsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccountsResponse struct{}"
	}

	return strings.Join([]string{"ListAccountsResponse", string(data)}, " ")
}
