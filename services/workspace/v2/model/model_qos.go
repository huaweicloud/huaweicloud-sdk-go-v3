package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Qos QOS信息。
type Qos struct {

	// iops，磁盘每秒进行读写的操作次数。
	Iops int32 `json:"iops"`

	// 吞吐量，磁盘每秒成功传送的数据量，即读取和写入的数据量，单位是MiB/s。
	Throughput int32 `json:"throughput"`
}

func (o Qos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Qos struct{}"
	}

	return strings.Join([]string{"Qos", string(data)}, " ")
}
