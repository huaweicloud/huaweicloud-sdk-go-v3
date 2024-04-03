package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchPushCertificateResponseBodyResults struct {

	// 推送区域名称，如cn-north-7。
	ProjectName *string `json:"project_name,omitempty"`

	// 目标证书ID。
	CertId *string `json:"cert_id,omitempty"`

	// 推送结果。
	Message *string `json:"message,omitempty"`
}

func (o BatchPushCertificateResponseBodyResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchPushCertificateResponseBodyResults struct{}"
	}

	return strings.Join([]string{"BatchPushCertificateResponseBodyResults", string(data)}, " ")
}
