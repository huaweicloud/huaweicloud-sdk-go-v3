/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 批量操作资源标签的请求体
type BatchCreatePublicipTagsRequestBody struct {
	// 标签列表
	Tags []ResourceTagOption `json:"tags"`
	// 操作标识  create：创建  action为create时，tag的value必选
	Action string `json:"action"`
}
