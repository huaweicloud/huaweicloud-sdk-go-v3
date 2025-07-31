package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkPolicyBody 网络策略内容
type NetworkPolicyBody struct {
}

func (o NetworkPolicyBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkPolicyBody struct{}"
	}

	return strings.Join([]string{"NetworkPolicyBody", string(data)}, " ")
}
