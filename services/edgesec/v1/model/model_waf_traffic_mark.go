package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WafTrafficMark 流量标识（用于攻击惩罚）
type WafTrafficMark struct {

	// 惩罚ip
	Sip *[]string `json:"sip,omitempty"`

	// cookie
	Cookie *string `json:"cookie,omitempty"`

	// 参数
	Params *string `json:"params,omitempty"`
}

func (o WafTrafficMark) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WafTrafficMark struct{}"
	}

	return strings.Join([]string{"WafTrafficMark", string(data)}, " ")
}
