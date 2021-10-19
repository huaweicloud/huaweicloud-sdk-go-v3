package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListCommonPoolsResponse struct {
	// 功能说明：公共池对象

	CommonPools *[]CommonPoolDict `json:"common_pools,omitempty"`
	// 本次请求的编号

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCommonPoolsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCommonPoolsResponse struct{}"
	}

	return strings.Join([]string{"ListCommonPoolsResponse", string(data)}, " ")
}
