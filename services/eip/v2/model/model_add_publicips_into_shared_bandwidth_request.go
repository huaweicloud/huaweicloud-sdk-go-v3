/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Request Object
type AddPublicipsIntoSharedBandwidthRequest struct {
	BandwidthId string `json:"bandwidth_id"`
	Body *AddPublicipsIntoSharedBandwidthRequestBody `json:"body,omitempty"`
}
