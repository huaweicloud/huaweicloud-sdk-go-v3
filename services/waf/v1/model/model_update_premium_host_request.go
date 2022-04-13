package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePremiumHostRequest struct {
	// 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 独享模式域名ID

	HostId string `json:"host_id"`

	Body *UpdatePremiumHostRequestBody `json:"body,omitempty"`
}

func (o UpdatePremiumHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePremiumHostRequest struct{}"
	}

	return strings.Join([]string{"UpdatePremiumHostRequest", string(data)}, " ")
}
