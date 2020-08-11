/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 创建tag对象的请求体
type CreatePublicipTagRequestBody struct {
	Tag *ResourceTagOption `json:"tag"`
}
