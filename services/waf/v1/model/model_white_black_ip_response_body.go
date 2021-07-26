package model

import (
	"encoding/json"

	"strings"
)

type WhiteBlackIpResponseBody struct {
	// 规则id

	Id *string `json:"id,omitempty"`
	// 黑白名单

	Ip *string `json:"ip,omitempty"`
	// 类型，0拦截，1放行

	White *int32 `json:"white,omitempty"`
	// 创建规则的时间戳

	Timestamp *int64 `json:"timestamp,omitempty"`
}

func (o WhiteBlackIpResponseBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "WhiteBlackIpResponseBody struct{}"
	}

	return strings.Join([]string{"WhiteBlackIpResponseBody", string(data)}, " ")
}
