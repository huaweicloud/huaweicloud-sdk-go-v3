/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// 批量创建带宽的请求体
type BatchCreateBandwidthRequestBody struct {
	Bandwidth *BatchCreateBandwidthOption `json:"bandwidth"`
}
