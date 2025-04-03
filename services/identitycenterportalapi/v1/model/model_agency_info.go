package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyInfo 包含委托或信任委托信息的对象
type AgencyInfo struct {

	// 分配给用户的账号的全局唯一标识符（ID）
	AccountId *string `json:"account_id,omitempty"`

	// 分配给用户的委托或信任委托的名称
	AgencyName *string `json:"agency_name,omitempty"`

	// 权限集名称
	PermissionSetName *string `json:"permission_set_name,omitempty"`

	// 委托或信任委托的统一资源名称（URN）
	AgencyUrn *string `json:"agency_urn,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`
}

func (o AgencyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyInfo struct{}"
	}

	return strings.Join([]string{"AgencyInfo", string(data)}, " ")
}
