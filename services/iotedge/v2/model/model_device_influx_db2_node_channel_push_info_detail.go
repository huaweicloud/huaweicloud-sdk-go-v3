package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceInfluxDb2NodeChannelPushInfoDetail 创建设备数据推送InfluxDB2
type DeviceInfluxDb2NodeChannelPushInfoDetail struct {

	// 一组用户的工作空间，一组用户下可以创建多个bucket
	Organization *string `json:"organization,omitempty"`

	// 数据存储的地方，结合了数据库和存储周期的概念
	Bucket *string `json:"bucket,omitempty"`
}

func (o DeviceInfluxDb2NodeChannelPushInfoDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceInfluxDb2NodeChannelPushInfoDetail struct{}"
	}

	return strings.Join([]string{"DeviceInfluxDb2NodeChannelPushInfoDetail", string(data)}, " ")
}
