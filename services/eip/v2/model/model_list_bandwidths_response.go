/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Response Object
type ListBandwidthsResponse struct {
	// 带宽列表对象
	Bandwidths []BandwidthResp `json:"bandwidths,omitempty"`
}
