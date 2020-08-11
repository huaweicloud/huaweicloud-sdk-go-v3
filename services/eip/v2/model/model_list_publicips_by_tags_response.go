/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Response Object
type ListPublicipsByTagsResponse struct {
	// resource对象列表
	Resources []ListResourceResp `json:"resources,omitempty"`
	// 总记录数
	TotalCount int32 `json:"total_count,omitempty"`
}
