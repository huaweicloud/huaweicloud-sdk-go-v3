package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgreeTenantAuthorizationV2Req struct {

	// 授权详情列表
	AuthDetailList *[]TenantAgreeAuthDetailV2 `json:"auth_detail_list,omitempty" xml:"auth_detail_list"`

	// 授权生效时间
	AuthEffectiveTime *int64 `json:"auth_effective_time,omitempty" xml:"auth_effective_time"`

	// 授权到期时间
	AuthExpireTime *int64 `json:"auth_expire_time,omitempty" xml:"auth_expire_time"`

	// 组id
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 委托id
	AgencyId *string `json:"agency_id,omitempty" xml:"agency_id"`
}

func (o AgreeTenantAuthorizationV2Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgreeTenantAuthorizationV2Req struct{}"
	}

	return strings.Join([]string{"AgreeTenantAuthorizationV2Req", string(data)}, " ")
}
