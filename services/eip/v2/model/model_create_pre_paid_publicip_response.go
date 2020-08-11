/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Response Object
type CreatePrePaidPublicipResponse struct {
	Publicip *PublicipCreateResp `json:"publicip,omitempty"`
	// 订单号（预付费场景返回该字段）
	OrderId string `json:"order_id,omitempty"`
	// 弹性公网IP的ID（预付费场景返回该字段）
	PublicipId string `json:"publicip_id,omitempty"`
}
