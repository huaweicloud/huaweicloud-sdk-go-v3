package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartVpecpReq struct {

	// 内网域名。
	EndpointWithDnsName string `json:"endpointWithDnsName"`
}

func (o StartVpecpReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartVpecpReq struct{}"
	}

	return strings.Join([]string{"StartVpecpReq", string(data)}, " ")
}
