/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Response Object
type CreatePublicipResponse struct {
	Publicip *PublicipCreateResp `json:"publicip,omitempty"`
}
