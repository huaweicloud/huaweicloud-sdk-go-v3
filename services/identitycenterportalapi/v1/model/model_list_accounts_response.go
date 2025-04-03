package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccountsResponse Response Object
type ListAccountsResponse struct {

	// 满足查询条件的账号对象列表
	AccountList *[]AccountInfo `json:"account_list,omitempty"`

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
