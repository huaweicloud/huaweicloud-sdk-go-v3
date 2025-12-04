package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Scanner 漏洞扫描器
type Scanner struct {

	// 扫描器的名称
	Name *string `json:"name,omitempty"`

	// 扫描器的提供商
	Vendor *string `json:"vendor,omitempty"`

	// 扫描器的版本号
	Version *string `json:"version,omitempty"`
}

func (o Scanner) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Scanner struct{}"
	}

	return strings.Join([]string{"Scanner", string(data)}, " ")
}
