/*
    * LiveAPI
    *
    * 直播服务源站所有接口
    *
*/

package model

// Request Object
type ListRecordConfigsRequest struct {
	Domain string `json:"domain"`
	AppName string `json:"app_name,omitempty"`
	StreamName string `json:"stream_name,omitempty"`
	Page int32 `json:"page,omitempty"`
	Size int32 `json:"size,omitempty"`
	RecordType string `json:"record_type,omitempty"`
}
