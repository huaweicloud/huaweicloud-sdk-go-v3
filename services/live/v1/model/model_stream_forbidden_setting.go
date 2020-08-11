/*
    * LiveAPI
    *
    * 直播服务源站所有接口
    *
*/

package model
import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
)

type StreamForbiddenSetting struct {
	// 直播播放域名或推流域名
	Domain string `json:"domain"`
	// 流应用名称
	AppName string `json:"app_name"`
	// 流名称
	StreamName string `json:"stream_name"`
	// 恢复流时间，格式：yyyy-mm-ddThh:mm:ssZ，UTC时间，不指定则永久禁播
	ResumeTime *sdktime.SdkTime `json:"resume_time,omitempty"`
}
