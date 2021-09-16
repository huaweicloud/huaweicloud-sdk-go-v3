package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListFirewallsRequest struct {
	// 每页返回的个数  取值范围：0~1000

	Limit *int32 `json:"limit,omitempty"`
	// 查询的偏移量。

	Offset *int32 `json:"offset,omitempty"`
	// 通过ID过滤网络ACL。

	Id *string `json:"id,omitempty"`
	// 通过name模糊匹配网络ACL。

	Name *string `json:"name,omitempty"`
}

func (o ListFirewallsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListFirewallsRequest struct{}"
	}

	return strings.Join([]string{"ListFirewallsRequest", string(data)}, " ")
}
