package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrafficMark 流量标识（用于攻击惩罚），WAF根据这些配置判断如何在Header中识别客户端IP（代理模式）、如何在Cookie中识别Session、如何在参数中识别User。
type TrafficMark struct {

	// IP标记，客户端最原始的IP地址的HTTP请求头字段。
	Sip *[]string `json:"sip,omitempty"`

	// Session标记，用于Cookie恶意请求的攻击惩罚功能。在选择Cookie拦截的攻击惩罚功能前，必须配置该标识
	Cookie *string `json:"cookie,omitempty"`

	// User标记，用于Params恶意请求的攻击惩罚功能。在选择Params拦截的攻击惩罚功能前，必须配置该标识。
	Params *string `json:"params,omitempty"`
}

func (o TrafficMark) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrafficMark struct{}"
	}

	return strings.Join([]string{"TrafficMark", string(data)}, " ")
}
