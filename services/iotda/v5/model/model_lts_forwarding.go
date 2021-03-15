package model

import (
	"encoding/json"

	"strings"
)

// LTS转发配置信息
type LtsForwarding struct {
	// 用于接收日志的日志组Id，通过调用LTS服务创建日志组接口获取(https://support.huaweicloud.com/api-lts/lts_api_0012.html)。
	LogGroupId string `json:"log_group_id"`
	// 用于接收日志的日志流Id，通过调用LTS服务创建日志流接口获取(https://support.huaweicloud.com/api-lts/lts_api_0016.html)。
	LogStreamId string `json:"log_stream_id"`
}

func (o LtsForwarding) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "LtsForwarding struct{}"
	}

	return strings.Join([]string{"LtsForwarding", string(data)}, " ")
}
