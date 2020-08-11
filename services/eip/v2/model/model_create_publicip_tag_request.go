/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Request Object
type CreatePublicipTagRequest struct {
	PublicipId string `json:"publicip_id"`
	Body *CreatePublicipTagRequestBody `json:"body,omitempty"`
}
