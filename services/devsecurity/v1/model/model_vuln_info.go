package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulnInfo struct {

	// 漏洞类型
	Category *string `json:"category,omitempty"`

	// 漏洞原理
	Principle *string `json:"principle,omitempty"`

	// 解决方案
	Solution *string `json:"solution,omitempty"`

	// 漏洞项列表
	VulnItems *[]VulnItem `json:"vuln_items,omitempty"`
}

func (o VulnInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulnInfo struct{}"
	}

	return strings.Join([]string{"VulnInfo", string(data)}, " ")
}
