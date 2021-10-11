package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListNaAuthorizedNodesRequest struct {
	// 北向数据接收端点ID

	NaId string `json:"na_id"`
	// 查询的起始位置，取值范围为非负整数，默认为0

	Offset *int32 `json:"offset,omitempty"`
	// 每页记录数，默认值为10，取值区间为1-1000

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListNaAuthorizedNodesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListNaAuthorizedNodesRequest struct{}"
	}

	return strings.Join([]string{"ListNaAuthorizedNodesRequest", string(data)}, " ")
}
