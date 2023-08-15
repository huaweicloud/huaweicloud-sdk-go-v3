package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Probe struct {
	Exec *Exec `json:"exec,omitempty"`

	HttpGet *HttpGet `json:"http_get,omitempty"`

	// 表示从工作负载启动后从多久开始探测，大于0且不大于3600的整数，默认为10
	InitialDelaySeconds *int64 `json:"initial_delay_seconds,omitempty"`

	// 表示探测超时时间，大于0且不大于3600的整数，默认为1
	TimeoutSeconds *int64 `json:"timeout_seconds,omitempty"`
}

func (o Probe) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Probe struct{}"
	}

	return strings.Join([]string{"Probe", string(data)}, " ")
}
