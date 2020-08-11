/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Request Object
type UpdatePrePaidBandwidthRequest struct {
	BandwidthId string `json:"bandwidth_id"`
	Body *UpdatePrePaidBandwidthRequestBody `json:"body,omitempty"`
}
