package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDomainRequest Request Object
type DeleteDomainRequest struct {

	// 防护域名id
	DomainId string `json:"domain_id"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id，默认为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o DeleteDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainRequest struct{}"
	}

	return strings.Join([]string{"DeleteDomainRequest", string(data)}, " ")
}
