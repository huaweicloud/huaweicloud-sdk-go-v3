package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateDomainMultiCertificatesRequest struct {
	// 当用户开启企业项目功能时，该参数生效，表示资源所属项目，不传表示查询默认项目。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *UpdateDomainMultiCertificatesRequestBody `json:"body,omitempty"`
}

func (o UpdateDomainMultiCertificatesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainMultiCertificatesRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainMultiCertificatesRequest", string(data)}, " ")
}
