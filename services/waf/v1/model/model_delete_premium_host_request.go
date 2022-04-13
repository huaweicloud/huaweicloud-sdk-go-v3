package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeletePremiumHostRequest struct {
	// 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 独享模式域名ID

	HostId string `json:"host_id"`
	// 是否保留规则

	KeepPolicy *bool `json:"keepPolicy,omitempty"`
}

func (o DeletePremiumHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePremiumHostRequest struct{}"
	}

	return strings.Join([]string{"DeletePremiumHostRequest", string(data)}, " ")
}
