/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Response Object
type ShowPublicipTagsResponse struct {
	// 标签列表
	Tags []ResourceTagResp `json:"tags,omitempty"`
}
