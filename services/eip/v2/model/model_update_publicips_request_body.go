/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 更新弹性公网IP的请求体
type UpdatePublicipsRequestBody struct {
	Publicip *UpdatePublicipOption `json:"publicip"`
}
