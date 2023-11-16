package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOpenSourceStrategyRequest Request Object
type ListOpenSourceStrategyRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 策略名称
	Name *string `json:"name,omitempty"`

	// 策略创建人名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 分页参数，默认15
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，默认0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListOpenSourceStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOpenSourceStrategyRequest struct{}"
	}

	return strings.Join([]string{"ListOpenSourceStrategyRequest", string(data)}, " ")
}
