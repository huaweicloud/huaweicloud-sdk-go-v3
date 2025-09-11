package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadHttpCertRequest Request Object
type DownloadHttpCertRequest struct {

	// 证书文件ID名称，格式为SMN-{RegionID}-{UUID}.pem
	CertificateId string `json:"certificate_id"`
}

func (o DownloadHttpCertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadHttpCertRequest struct{}"
	}

	return strings.Join([]string{"DownloadHttpCertRequest", string(data)}, " ")
}
