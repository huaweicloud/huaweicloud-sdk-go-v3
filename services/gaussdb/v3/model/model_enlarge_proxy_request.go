package model

import (
	"encoding/json"

	"strings"
)

// proxy节点扩容信息
type EnlargeProxyRequest struct {
	// proxy节点扩容。取值范围：1~30之间的整数，proxy节点的总数量小于等于32个。

	NodeNum int32 `json:"node_num"`
}

func (o EnlargeProxyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "EnlargeProxyRequest struct{}"
	}

	return strings.Join([]string{"EnlargeProxyRequest", string(data)}, " ")
}
