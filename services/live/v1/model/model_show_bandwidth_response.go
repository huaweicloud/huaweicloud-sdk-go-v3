/*
    * LiveAPI
    *
    * 直播服务源站所有接口
    *
*/

package model

// Response Object
type ShowBandwidthResponse struct {
	// 查询结果的总元素数量
	Total int32 `json:"total,omitempty"`
	// 带宽信息
	BandwidthInfo []BandwidthInfo `json:"bandwidth_info,omitempty"`
}
