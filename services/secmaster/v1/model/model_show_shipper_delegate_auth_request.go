package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowShipperDelegateAuthRequest Request Object
type ShowShipperDelegateAuthRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 账号ID
	DomainId string `json:"domain_id"`

	// 委托的名称。委托是由租户管理员在统一身份认证服务（Identity and Access Management，IAM）上创建的，可以为弹性云服务器提供访问云服务器的临时凭证。
	AgencyName string `json:"agency_name"`
}

func (o ShowShipperDelegateAuthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShipperDelegateAuthRequest struct{}"
	}

	return strings.Join([]string{"ShowShipperDelegateAuthRequest", string(data)}, " ")
}
