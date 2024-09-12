package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDomainsRequest Request Object
type DeleteDomainsRequest struct {

	// 防护域名id
	DomainId string `json:"domain_id"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o DeleteDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainsRequest struct{}"
	}

	return strings.Join([]string{"DeleteDomainsRequest", string(data)}, " ")
}
