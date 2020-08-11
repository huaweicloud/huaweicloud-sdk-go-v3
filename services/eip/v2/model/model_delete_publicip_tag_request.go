/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Request Object
type DeletePublicipTagRequest struct {
	PublicipId string `json:"publicip_id"`
	Key string `json:"key"`
}
