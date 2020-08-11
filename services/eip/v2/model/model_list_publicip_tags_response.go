/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Response Object
type ListPublicipTagsResponse struct {
	// 标签列表
	Tags []TagResp `json:"tags,omitempty"`
}
