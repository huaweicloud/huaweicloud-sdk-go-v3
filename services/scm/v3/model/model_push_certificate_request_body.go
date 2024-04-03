package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PushCertificateRequestBody struct {

	// 推送到的目标服务所在的区域。
	TargetProject string `json:"target_project"`

	// 证书推送的目标服务，当前仅支持：CDN、WAF、ELB。
	TargetService string `json:"target_service"`
}

func (o PushCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"PushCertificateRequestBody", string(data)}, " ")
}
