/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 共享带宽移除弹性公网IP的请求体
type RemovePublicipsFromSharedBandwidthRequestBody struct {
	Bandwidth *RemoveFromSharedBandwidthOption `json:"bandwidth"`
}
