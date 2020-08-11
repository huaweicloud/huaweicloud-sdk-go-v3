/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 标签
type TagResp struct {
	// 键，key不能为空。长度不超过36个字符。由英文字母、数字、下划线、中划线、中文字符组成。
	Key string `json:"key,omitempty"`
	// 值列表。
	Values []string `json:"values,omitempty"`
}
