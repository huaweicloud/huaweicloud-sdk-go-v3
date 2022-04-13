package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 流量标识（用于攻击惩罚）
type TrafficMark struct {
	// 惩罚ip

	Sip *[]string `json:"sip,omitempty"`
	// cookie

	Cookie *string `json:"cookie,omitempty"`
	// 参数

	Params *string `json:"params,omitempty"`
}

func (o TrafficMark) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrafficMark struct{}"
	}

	return strings.Join([]string{"TrafficMark", string(data)}, " ")
}
