package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPartnersCasesPrivilegeResponse Response Object
type ShowPartnersCasesPrivilegeResponse struct {

	// 是否有权限
	HasPrivilege   *bool `json:"has_privilege,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowPartnersCasesPrivilegeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPartnersCasesPrivilegeResponse struct{}"
	}

	return strings.Join([]string{"ShowPartnersCasesPrivilegeResponse", string(data)}, " ")
}
