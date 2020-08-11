/*
    * LiveAPI
    *
    * 直播服务源站所有接口
    *
*/

package model

// Response Object
type ListRecordConfigsResponse struct {
	// 查询结果的总元素数量
	Total int32 `json:"total,omitempty"`
	// 录制配置数组
	RecordConfig []RecordConfigInfo `json:"record_config,omitempty"`
}
