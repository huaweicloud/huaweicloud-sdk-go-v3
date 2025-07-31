package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccountsResponse Response Object
type ListAccountsResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// 事件列表详情
	DataList *[]AccountResponseInfo `json:"data_list,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAccountsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccountsResponse struct{}"
	}

	return strings.Join([]string{"ListAccountsResponse", string(data)}, " ")
}
