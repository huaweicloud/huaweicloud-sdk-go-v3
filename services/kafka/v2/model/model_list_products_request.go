package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListProductsRequest struct {
	// 消息引擎的类型。当前只支持kafka类型。

	Engine string `json:"engine"`
}

func (o ListProductsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProductsRequest struct{}"
	}

	return strings.Join([]string{"ListProductsRequest", string(data)}, " ")
}
