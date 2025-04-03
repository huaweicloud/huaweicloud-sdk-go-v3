package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceProviderConfigRequest Request Object
type ServiceProviderConfigRequest struct {

	// 承载令牌
	Authorization string `json:"Authorization"`

	// 租户的全局唯一标识符（ID）
	TenantId string `json:"tenant_id"`
}

func (o ServiceProviderConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceProviderConfigRequest struct{}"
	}

	return strings.Join([]string{"ServiceProviderConfigRequest", string(data)}, " ")
}
