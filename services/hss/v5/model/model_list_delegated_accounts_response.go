package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDelegatedAccountsResponse Response Object
type ListDelegatedAccountsResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释** 已委托账号列表信息 **取值范围** 数组长度0-10000
	DataList *[]DelegatedAccountResponseInfo `json:"data_list,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDelegatedAccountsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDelegatedAccountsResponse struct{}"
	}

	return strings.Join([]string{"ListDelegatedAccountsResponse", string(data)}, " ")
}
