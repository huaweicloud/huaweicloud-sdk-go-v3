/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 创建共享带宽请求体
type CreateSharedBandwidhRequestBody struct {
	Bandwidth *CreateSharedBandwidthOption `json:"bandwidth"`
}
