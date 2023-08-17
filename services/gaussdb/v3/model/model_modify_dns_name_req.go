package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyDnsNameReq 修改内网域名
type ModifyDnsNameReq struct {

	// 新域名的前缀，取值范围如下：8~63个字符之间，可以包含小写字母、数字，不能包含其他特殊字符。
	DnsName string `json:"dns_name"`
}

func (o ModifyDnsNameReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyDnsNameReq struct{}"
	}

	return strings.Join([]string{"ModifyDnsNameReq", string(data)}, " ")
}
