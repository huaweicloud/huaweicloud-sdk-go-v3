package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BlackWhiteIpListRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`

	// 规则类型。black - 黑名单， white - 白名单
	Type string `json:"type"`

	// ip列表
	Ips []string `json:"ips"`
}

func (o BlackWhiteIpListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BlackWhiteIpListRequest struct{}"
	}

	return strings.Join([]string{"BlackWhiteIpListRequest", string(data)}, " ")
}
