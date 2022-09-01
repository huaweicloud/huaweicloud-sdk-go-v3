package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type NovaSecurityGroupCommonGroup struct {

	// 对端安全组的名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 对端安全组所属租户的租户ID
	TenantId *string `json:"tenant_id,omitempty" xml:"tenant_id"`
}

func (o NovaSecurityGroupCommonGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaSecurityGroupCommonGroup struct{}"
	}

	return strings.Join([]string{"NovaSecurityGroupCommonGroup", string(data)}, " ")
}
