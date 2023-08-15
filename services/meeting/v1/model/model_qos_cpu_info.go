package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QosCpuInfo 会议与会者CPU QoS数据，当qosType = cpu 时有效。
type QosCpuInfo struct {

	// App最大CPU使用率。
	ClientCpuMax *[]QosDataElement `json:"client_cpu_max,omitempty"`

	// 系统最大CPU使用率。
	SystemCpuMax *[]QosDataElement `json:"system_cpu_max,omitempty"`
}

func (o QosCpuInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QosCpuInfo struct{}"
	}

	return strings.Join([]string{"QosCpuInfo", string(data)}, " ")
}
