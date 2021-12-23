package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgreeTenantAuthorizationV2Req struct {
	// 授权详情列表

	AuthDetailList *[]TenantAgreeAuthDetailV2 `json:"auth_detail_list,omitempty"`
	// 授权生效时间

	AuthEffectiveTime *int64 `json:"auth_effective_time,omitempty"`
	// 授权到期时间

	AuthExpireTime *int64 `json:"auth_expire_time,omitempty"`
	// 组id

	GroupId *string `json:"group_id,omitempty"`
	// 委托id

	AgencyId *string `json:"agency_id,omitempty"`
}

func (o AgreeTenantAuthorizationV2Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgreeTenantAuthorizationV2Req struct{}"
	}

	return strings.Join([]string{"AgreeTenantAuthorizationV2Req", string(data)}, " ")
}
