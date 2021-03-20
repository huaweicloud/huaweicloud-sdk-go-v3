package model

import (
	"encoding/json"

	"strings"
)

type ModifyDnsNameRequestBody struct {
	// 新域名的前缀，校验规则是^[0-9a-zA-Z]{8,64}$

	DnsName string `json:"dns_name"`
}

func (o ModifyDnsNameRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ModifyDnsNameRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyDnsNameRequestBody", string(data)}, " ")
}
