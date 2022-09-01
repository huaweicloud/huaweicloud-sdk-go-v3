package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 探针配置
type ProbeDetail struct {

	// 执行探测的命令行命令，长度1-10240内的字符串
	ExecCommand *string `json:"exec_command,omitempty" xml:"exec_command"`

	HttpGet *HttpGetDetail `json:"http_get,omitempty" xml:"http_get"`

	// 表示从工作负载启动后从多久开始探测，大于0且不大于3600的整数，默认为10
	InitialDelaySeconds *int32 `json:"initial_delay_seconds,omitempty" xml:"initial_delay_seconds"`

	// 表示探测超时时间，大于0且不大于3600的整数，默认为1
	TimeoutSeconds *int32 `json:"timeout_seconds,omitempty" xml:"timeout_seconds"`
}

func (o ProbeDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProbeDetail struct{}"
	}

	return strings.Join([]string{"ProbeDetail", string(data)}, " ")
}
