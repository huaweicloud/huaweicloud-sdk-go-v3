package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeviceInfluxDb2PushInfo 创建设备数据推送MQTT
type DeviceInfluxDb2PushInfo struct {

	// 一组用户的工作空间，一组用户下可以创建多个bucket
	Organization string `json:"organization"`

	// 数据存储的地方，结合了数据库和存储周期的概念
	Bucket string `json:"bucket"`

	// 数据格式转换类型
	Format string `json:"format"`
}

func (o DeviceInfluxDb2PushInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeviceInfluxDb2PushInfo struct{}"
	}

	return strings.Join([]string{"DeviceInfluxDb2PushInfo", string(data)}, " ")
}
