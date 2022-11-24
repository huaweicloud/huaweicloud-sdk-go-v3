package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetPrivateZoneProxyPatternRequestBody struct {

	// 内网Zone的子域名递归解析代理模式。  取值范围：  AUTHORITY：当前Zone不进行递归解析 RECURSIVE：开启递归解析代理
	ProxyPattern string `json:"proxy_pattern"`
}

func (o SetPrivateZoneProxyPatternRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetPrivateZoneProxyPatternRequestBody struct{}"
	}

	return strings.Join([]string{"SetPrivateZoneProxyPatternRequestBody", string(data)}, " ")
}
