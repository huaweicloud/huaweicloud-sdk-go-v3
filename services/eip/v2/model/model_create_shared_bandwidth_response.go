/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Response Object
type CreateSharedBandwidthResponse struct {
	Bandwidth *BandwidthResp `json:"bandwidth,omitempty"`
}
