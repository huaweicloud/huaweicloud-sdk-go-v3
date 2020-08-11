/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 将弹性公网IP插入共享带宽的请求体
type AddPublicipsIntoSharedBandwidthRequestBody struct {
	Bandwidth *AddPublicipsIntoSharedBandwidthOption `json:"bandwidth"`
}
