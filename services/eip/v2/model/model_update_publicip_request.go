/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Request Object
type UpdatePublicipRequest struct {
	PublicipId string `json:"publicip_id"`
	Body *UpdatePublicipsRequestBody `json:"body,omitempty"`
}
