package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadHttpSignCertRequest Request Object
type DownloadHttpSignCertRequest struct {

	// 证书文件ID名称，格式为SMN-{RegionID}-{UUID}.pem
	CertificateId string `json:"certificate_id"`
}

func (o DownloadHttpSignCertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadHttpSignCertRequest struct{}"
	}

	return strings.Join([]string{"DownloadHttpSignCertRequest", string(data)}, " ")
}
