package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPolicyResponse Response Object
type ShowPolicyResponse struct {

	// 保护类型
	ProtectionType *string `json:"protection_type,omitempty"`

	// 策略ID
	Id *string `json:"id,omitempty"`

	// 是否启用策略
	Enabled *bool `json:"enabled,omitempty"`

	// 是否启用加密
	Cyber *string `json:"cyber,omitempty"`

	// 区域名称
	Name *string `json:"name,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 存储类型，默认obs
	StorageType *string `json:"storage_type,omitempty"`

	// 策略类型
	PolicyType *string `json:"policy_type,omitempty"`

	// 保护服务类型
	ServiceType *string `json:"service_type,omitempty"`

	CbrPolicyDetail *CbrPolicyDetail `json:"cbr_policy_detail,omitempty"`

	DbPolicyDetail *DbPolicyDetail `json:"db_policy_detail,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowPolicyResponse", string(data)}, " ")
}
