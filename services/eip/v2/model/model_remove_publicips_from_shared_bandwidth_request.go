/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Request Object
type RemovePublicipsFromSharedBandwidthRequest struct {
	BandwidthId string `json:"bandwidth_id"`
	Body *RemovePublicipsFromSharedBandwidthRequestBody `json:"body,omitempty"`
}
