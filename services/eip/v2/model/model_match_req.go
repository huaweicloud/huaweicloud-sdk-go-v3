/*
 * EIP
 *
 * 云服务接口
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// 搜索字段
type MatchReq struct {
	// 键。当前仅限定为resource_name
	Key MatchReqKey `json:"key"`
	// 值。每个值最大长度255个unicode字符。
	Value string `json:"value"`
}

func (o MatchReq) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"MatchReq", string(data)}, " ")
}

type MatchReqKey struct {
	value string
}

type MatchReqKeyEnum struct {
	RESOURCE_NAME MatchReqKey
}

func GetMatchReqKeyEnum() MatchReqKeyEnum {
	return MatchReqKeyEnum{
		RESOURCE_NAME: MatchReqKey{
			value: "resource_name",
		},
	}
}

func (c MatchReqKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *MatchReqKey) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
