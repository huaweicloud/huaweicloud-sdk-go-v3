package model

import (
	"encoding/json"

	"strings"
)

//
type PoolRef struct {
	// 功能描述：后端服务器组ID。

	Id string `json:"id"`
}

func (o PoolRef) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PoolRef struct{}"
	}

	return strings.Join([]string{"PoolRef", string(data)}, " ")
}
