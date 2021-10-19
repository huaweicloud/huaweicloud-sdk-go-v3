package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListPublicBorderGroupsRequest struct {
	// 显示，形式为\"fields=id&fields=name&...\"  支持字段：publicip_pools/public_border_group

	Fields *string `json:"fields,omitempty"`
}

func (o ListPublicBorderGroupsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPublicBorderGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListPublicBorderGroupsRequest", string(data)}, " ")
}
