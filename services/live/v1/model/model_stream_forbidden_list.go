/*
    * LiveAPI
    *
    * 直播服务源站所有接口
    *
*/

package model

type StreamForbiddenList struct {
	// 流应用名称
	AppName string `json:"app_name"`
	// 流名称
	StreamName string `json:"stream_name"`
	// 恢复流时间，格式：yyyy-mm-ddThh:mm:ssZ，UTC时间，不指定则永久禁播
	ResumeTime string `json:"resume_time,omitempty"`
}
