/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Response Object
type ShowBandwidthResponse struct {
	Bandwidth *BandwidthResp `json:"bandwidth,omitempty"`
}
