/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Request Object
type BatchCreateSharedBandwidthsRequest struct {
	Body *BatchCreateBandwidthRequestBody `json:"body,omitempty"`
}
