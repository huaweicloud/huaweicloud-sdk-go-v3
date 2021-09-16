package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListSubnetsRequest struct {
	// 虚拟私有云ID。

	VpcId *string `json:"vpc_id,omitempty"`
	// 查询返回边缘子网列表数量。取值范围：0~1000。

	Limit *int32 `json:"limit,omitempty"`
	// 查询的偏移量。

	Offset *int32 `json:"offset,omitempty"`
	// 站点ID。

	SiteId *string `json:"site_id,omitempty"`
}

func (o ListSubnetsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSubnetsRequest struct{}"
	}

	return strings.Join([]string{"ListSubnetsRequest", string(data)}, " ")
}
