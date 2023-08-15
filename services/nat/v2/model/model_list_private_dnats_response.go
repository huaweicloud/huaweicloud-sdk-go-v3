package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrivateDnatsResponse Response Object
type ListPrivateDnatsResponse struct {

	// 查询DNAT规则列表的响应体。
	DnatRules *[]PrivateDnat `json:"dnat_rules,omitempty"`

	// 请求ID。
	RequestId *string `json:"request_id,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPrivateDnatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateDnatsResponse struct{}"
	}

	return strings.Join([]string{"ListPrivateDnatsResponse", string(data)}, " ")
}
