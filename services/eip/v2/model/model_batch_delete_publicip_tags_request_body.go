/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 批量操作资源标签的请求体
type BatchDeletePublicipTagsRequestBody struct {
	// 标签列表
	Tags []ResourceTagOption `json:"tags"`
	// 操作标识  delete：删除  action为delete时，value可选
	Action string `json:"action"`
}
