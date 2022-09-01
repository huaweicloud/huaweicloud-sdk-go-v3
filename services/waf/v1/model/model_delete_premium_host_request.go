package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeletePremiumHostRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 独享模式域名ID
	HostId string `json:"host_id" xml:"host_id"`

	// 是否保留规则。false表示不保留该域名的防护策略；true表示保留该域名的防护策略。当要删除的防护域名的防护策略防护多个防护域名时，该参数不传。
	KeepPolicy *bool `json:"keepPolicy,omitempty" xml:"keepPolicy"`
}

func (o DeletePremiumHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePremiumHostRequest struct{}"
	}

	return strings.Join([]string{"DeletePremiumHostRequest", string(data)}, " ")
}
