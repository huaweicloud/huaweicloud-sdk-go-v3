package model

import (
	"encoding/json"

	"strings"
)

// 绑定证书和域名的请求体
type ApplyCertificateToHostRequestBody struct {
	// 云模式HTTPS域名ID

	CloudHostIds *[]string `json:"cloud_host_ids,omitempty"`
	// 独享模式HTTPS域名ID

	PremiumHostIds *[]string `json:"premium_host_ids,omitempty"`
}

func (o ApplyCertificateToHostRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ApplyCertificateToHostRequestBody struct{}"
	}

	return strings.Join([]string{"ApplyCertificateToHostRequestBody", string(data)}, " ")
}
