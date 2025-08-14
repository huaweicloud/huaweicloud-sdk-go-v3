package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationInstanceCertificatesResponse Response Object
type ListApplicationInstanceCertificatesResponse struct {

	// 应用程序证书列表
	ApplicationInstanceCertificates *[]CertificateDto `json:"application_instance_certificates,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListApplicationInstanceCertificatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationInstanceCertificatesResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationInstanceCertificatesResponse", string(data)}, " ")
}
