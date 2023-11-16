package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOpenSourceStrategyRequest Request Object
type CreateOpenSourceStrategyRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	Body *TenantOpenSourceRuleSetCreateVo `json:"body,omitempty"`
}

func (o CreateOpenSourceStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOpenSourceStrategyRequest struct{}"
	}

	return strings.Join([]string{"CreateOpenSourceStrategyRequest", string(data)}, " ")
}
