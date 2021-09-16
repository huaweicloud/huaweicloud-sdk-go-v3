package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListVpcsRequest struct {
	// 查询返回虚拟私有云列表数量。

	Limit *int32 `json:"limit,omitempty"`
	// 查询的偏移量。

	Offset *int32 `json:"offset,omitempty"`
	// 通过ID查询

	Id *string `json:"id,omitempty"`
	// 通过name查询

	Name *string `json:"name,omitempty"`
}

func (o ListVpcsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListVpcsRequest struct{}"
	}

	return strings.Join([]string{"ListVpcsRequest", string(data)}, " ")
}
