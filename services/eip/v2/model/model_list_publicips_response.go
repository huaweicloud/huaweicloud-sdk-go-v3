/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Response Object
type ListPublicipsResponse struct {
	// 弹性公网IP对象
	Publicips []PublicipShowResp `json:"publicips,omitempty"`
}
