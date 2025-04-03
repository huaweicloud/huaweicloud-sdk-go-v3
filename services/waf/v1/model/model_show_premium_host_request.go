package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPremiumHostRequest Request Object
type ShowPremiumHostRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 独享模式域名ID
	HostId string `json:"host_id"`
}

func (o ShowPremiumHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPremiumHostRequest struct{}"
	}

	return strings.Join([]string{"ShowPremiumHostRequest", string(data)}, " ")
}
