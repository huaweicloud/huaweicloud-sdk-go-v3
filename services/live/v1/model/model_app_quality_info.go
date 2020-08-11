/*
    * LiveAPI
    *
    * 直播服务源站所有接口
    *
*/

package model

type AppQualityInfo struct {
	// 应用名称
	AppName string `json:"app_name,omitempty"`
	// 视频质量信息
	QualityInfo []QualityInfo `json:"quality_info,omitempty"`
}
