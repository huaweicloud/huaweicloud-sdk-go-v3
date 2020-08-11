/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Response Object
type BatchCreateSharedBandwidthsResponse struct {
	// 批创的带宽对象的列表
	Bandwidths []BatchBandwidthResp `json:"bandwidths,omitempty"`
}
