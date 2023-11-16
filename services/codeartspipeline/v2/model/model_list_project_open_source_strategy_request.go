package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectOpenSourceStrategyRequest Request Object
type ListProjectOpenSourceStrategyRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 策略名称
	Name *string `json:"name,omitempty"`

	// 策略创建人名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 是否包含当前项目所属租户的策略
	IncludeTenantRuleSet bool `json:"include_tenant_rule_set"`

	// 分页参数，默认15
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，默认0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListProjectOpenSourceStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectOpenSourceStrategyRequest struct{}"
	}

	return strings.Join([]string{"ListProjectOpenSourceStrategyRequest", string(data)}, " ")
}
