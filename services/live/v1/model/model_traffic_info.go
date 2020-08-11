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

type TrafficInfo struct {
	// 采样周期内的总流量，单位：byte
	Traffic int32 `json:"traffic"`
	// 流量数据采样周期起始时刻，UTC时间，格式：yyyy-MM-ddTHH:mm:ssZ
	Timestamp *sdktime.SdkTime `json:"timestamp"`
}
