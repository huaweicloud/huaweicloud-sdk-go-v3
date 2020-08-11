/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Request Object
type UpdateBandwidthRequest struct {
	BandwidthId string `json:"bandwidth_id"`
	Body *UpdateBandwidthRequestBody `json:"body,omitempty"`
}
