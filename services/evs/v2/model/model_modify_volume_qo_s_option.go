package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyVolumeQoSOption
type ModifyVolumeQoSOption struct {

	// 修改后的云硬盘iops，只支持GPSSD2、ESSD2类型的云硬盘。  说明： 了解GPSSD2、ESSD2类型的iops大小范围，请参见 [云硬盘类型及性能介绍里面的云硬盘性能数据表](https://support.huaweicloud.com/productdesc-evs/zh-cn_topic_0044524691.html)。
	Iops int32 `json:"iops"`

	// 修改后的云硬盘吞吐量，单位是MiB/s，GPSSD2类型云盘必须填写，其他类型不能填写。  说明： 了解GPSSD2类型的吞吐量大小范围，请参见 [云硬盘类型及性能介绍里面的云硬盘性能数据表](https://support.huaweicloud.com/productdesc-evs/zh-cn_topic_0044524691.html)。
	Throughput *int32 `json:"throughput,omitempty"`
}

func (o ModifyVolumeQoSOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyVolumeQoSOption struct{}"
	}

	return strings.Join([]string{"ModifyVolumeQoSOption", string(data)}, " ")
}
