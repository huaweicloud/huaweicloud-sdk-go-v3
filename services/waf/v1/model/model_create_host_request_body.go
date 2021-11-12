package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateHostRequestBody struct {
	// 域名

	Hostname *string `json:"hostname,omitempty"`
	// 防护域名初始绑定的策略ID

	Policyid *string `json:"policyid,omitempty"`
	// 源站信息

	Server *[]CloudWafServer `json:"server,omitempty"`
	// 证书id

	Certificateid *string `json:"certificateid,omitempty"`
	// 证书名

	Certificatename *string `json:"certificatename,omitempty"`
	// 是否使用代理

	Proxy *bool `json:"proxy,omitempty"`
	// 域名描述

	Description *string `json:"description,omitempty"`
}

func (o CreateHostRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHostRequestBody", string(data)}, " ")
}
