/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 标签
type ResourceTagResp struct {
	// 键。同一资源的key值不能重复。
	Key string `json:"key,omitempty"`
	// 值列表。
	Value string `json:"value,omitempty"`
}
