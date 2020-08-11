/*
    * LiveAPI
    *
    * 直播服务源站所有接口
    *
*/

package model

type RecordConfigInfo struct {
	// 直播播放域名
	Domain string `json:"domain"`
	// 应用名称。 默认为“live”，若您需要自定义应用名称，请先提交工单申请。 
	AppName string `json:"app_name"`
	// 秒，周期录制时常，最小15分钟，最大6小时，默认1小时
	RecordDuration int32 `json:"record_duration,omitempty"`
	// 录制格式flv，hls，mp4，默认flv（目前仅支持flv）
	RecordFormat string `json:"record_format,omitempty"`
	// 录制类型，configer_record，默认configer_record。表示录制配置好后，只要有流就录制
	RecordType string `json:"record_type,omitempty"`
	// 录制位置vod， 默认vod（目前暂只支持vod）
	RecordLocation string `json:"record_location,omitempty"`
	// 录制文件前缀， DomainName，AppName，StreamName必须，默认{DomainName}/{AppName}/{StreamName}/{StartTime}-{EndTime}
	RecordPrefix string `json:"record_prefix,omitempty"`
	ObsAddr *ObsFileAddr `json:"obs_addr,omitempty"`
	// 录制配置创建时间，样例2020-02-14T10:45:16.947+08:00
	CreateTime string `json:"create_time,omitempty"`
}
