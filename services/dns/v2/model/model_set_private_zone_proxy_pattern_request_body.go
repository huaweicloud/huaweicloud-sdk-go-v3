package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetPrivateZoneProxyPatternRequestBody struct {

	// 内网域名的子域名递归解析代理模式。  取值范围：  AUTHORITY：当前域名未开启递归解析代理 RECURSIVE：当前域名已开启递归解析代理
	ProxyPattern string `json:"proxy_pattern"`
}

func (o SetPrivateZoneProxyPatternRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetPrivateZoneProxyPatternRequestBody struct{}"
	}

	return strings.Join([]string{"SetPrivateZoneProxyPatternRequestBody", string(data)}, " ")
}
