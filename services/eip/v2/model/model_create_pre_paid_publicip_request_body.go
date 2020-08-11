/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 创建包周期的弹性公网IP
type CreatePrePaidPublicipRequestBody struct {
	Publicip *CreatePrePaidPublicipOption `json:"publicip"`
	Bandwidth *CreatePublicipBandwidthOption `json:"bandwidth"`
	ExtendParam *CreatePrePaidPublicipExtendParamOption `json:"extendParam,omitempty"`
}
