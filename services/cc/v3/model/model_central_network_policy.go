package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetworkPolicy 中心网络策略详情。
type CentralNetworkPolicy struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 实例所属账号ID。
	DomainId string `json:"domain_id"`

	State *CentralNetworkPolicyStateEnum `json:"state"`

	// 中心网络ID。
	CentralNetworkId string `json:"central_network_id"`

	DocumentTemplateVersion *DocumentTemplateVersionEnum `json:"document_template_version"`

	// 是否被应用。
	IsApplied bool `json:"is_applied"`

	// 中心网络策略的版本。
	Version int32 `json:"version"`

	Document *CentralNetworkPolicyDocument `json:"document"`
}

func (o CentralNetworkPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkPolicy struct{}"
	}

	return strings.Join([]string{"CentralNetworkPolicy", string(data)}, " ")
}
