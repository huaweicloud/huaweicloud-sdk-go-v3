/*
    * LiveAPI
    *
    * 直播服务源站所有接口
    *
*/

package model

type BandwidthInfo struct {
	// 带宽峰值，单位：bps
	BwBps int32 `json:"bw_bps"`
	// 带宽数据采样周期起始时刻，UTC时间，格式：yyyy-MM-ddTHH:mm:ssZ
	Timestamp string `json:"timestamp"`
}
