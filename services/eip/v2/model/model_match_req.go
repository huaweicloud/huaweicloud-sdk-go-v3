/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 搜索字段
type MatchReq struct {
	// 键。当前仅限定为resource_name
	Key string `json:"key"`
	// 值。每个值最大长度255个unicode字符。
	Value string `json:"value"`
}
