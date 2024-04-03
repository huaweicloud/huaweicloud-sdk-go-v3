package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchPushCertificateRequestBody struct {

	// 推送到的目标服务所在的区域，CDN、ELB、WAF。
	TargetProjects []string `json:"target_projects"`

	// 证书推送的目标服务，当前仅支持：CDN、WAF、ELB。
	TargetService string `json:"target_service"`
}

func (o BatchPushCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchPushCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"BatchPushCertificateRequestBody", string(data)}, " ")
}
