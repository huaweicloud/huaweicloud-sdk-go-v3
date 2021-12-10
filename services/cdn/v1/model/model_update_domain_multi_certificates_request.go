package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateDomainMultiCertificatesRequest struct {
	// 当用户开启企业项目功能时，该参数生效，表示查询资源所属项目，不传表示查询默认项目。注意：当使用子账号调用接口时，该参数必传。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *UpdateDomainMultiCertificatesRequestBody `json:"body,omitempty"`
}

func (o UpdateDomainMultiCertificatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainMultiCertificatesRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainMultiCertificatesRequest", string(data)}, " ")
}
