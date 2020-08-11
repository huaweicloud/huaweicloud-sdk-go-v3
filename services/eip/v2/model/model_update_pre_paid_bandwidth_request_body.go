/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 更新带宽的请求体
type UpdatePrePaidBandwidthRequestBody struct {
	Bandwidth *UpdatePrePaidBandwidthOption `json:"bandwidth"`
	ExtendParam *UpdatePrePaidBandwidthExtendParamOption `json:"extendParam,omitempty"`
}
