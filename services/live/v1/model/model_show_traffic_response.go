/*
    * LiveAPI
    *
    * 直播服务源站所有接口
    *
*/

package model

// Response Object
type ShowTrafficResponse struct {
	// 查询结果的总元素数量
	Total int32 `json:"total,omitempty"`
	// 流量信息
	TrafficInfo []TrafficInfo `json:"traffic_info,omitempty"`
}
