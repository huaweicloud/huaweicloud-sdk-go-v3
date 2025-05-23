package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertificateRequest Request Object
type ShowCertificateRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// https证书id，您可以通过调用查询证书列表（ListCertificates）接口获取证书id
	CertificateId string `json:"certificate_id"`
}

func (o ShowCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificateRequest struct{}"
	}

	return strings.Join([]string{"ShowCertificateRequest", string(data)}, " ")
}
