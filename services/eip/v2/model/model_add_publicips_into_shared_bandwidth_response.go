/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Response Object
type AddPublicipsIntoSharedBandwidthResponse struct {
	Bandwidth *BandwidthRespInsert `json:"bandwidth,omitempty"`
}
