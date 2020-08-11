/*
    * LiveAPI
    *
    * 直播服务源站所有接口
    *
*/

package model

type StreamTranscodingTemplate struct {
	// 播放域名
	Domain string `json:"domain"`
	// 应用名称。 默认为“live”，若您需要自定义应用名称，请先提交工单申请。 
	AppName string `json:"app_name"`
	// 视频质量信息
	QualityInfo []QualityInfo `json:"quality_info"`
}
