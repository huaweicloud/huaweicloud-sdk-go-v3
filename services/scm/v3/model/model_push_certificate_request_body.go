package model

import (
	"encoding/json"

	"strings"
)

type PushCertificateRequestBody struct {
	// 推送到的目标服务所在的区域。
	TargetProject *string `json:"target_project,omitempty"`
	// 证书推送的目标服务，当前仅支持：CDN、WAF、ELB。
	TargetService *string `json:"target_service,omitempty"`
}

func (o PushCertificateRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PushCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"PushCertificateRequestBody", string(data)}, " ")
}
