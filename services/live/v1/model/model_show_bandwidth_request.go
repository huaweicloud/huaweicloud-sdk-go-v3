/*
    * LiveAPI
    *
    * 直播服务源站所有接口
    *
*/

package model

// Request Object
type ShowBandwidthRequest struct {
	Domain string `json:"domain,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	EndTime string `json:"end_time,omitempty"`
	Step int32 `json:"step,omitempty"`
}
