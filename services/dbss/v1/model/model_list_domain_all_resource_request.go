package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainAllResourceRequest Request Object
type ListDomainAllResourceRequest struct {

	// **参数解释**： 账户ID。 通过调用IAM服务[查询IAM用户详情]接口获取 **约束限制**： 不涉及 **取值范围**： 以IAM服务[查询IAM用户详情]接口值为准。 **默认取值**： 不涉及
	DomainId string `json:"domain_id"`

	// **参数解释**：  资源类型 **约束限制**： 不涉及 **取值范围**：  - cloudservers: 审计  - dbEncrypt: 加密  - dbOm: 运维 **默认取值**： 不涉及
	ResourceType string `json:"resource_type"`

	// **参数解释**：  以当前region id为准 **约束限制**： 不涉及 **取值范围**： 以当前region id为准 **默认取值**： 以当前region id为准
	RegionId string `json:"region_id"`

	// **参数解释**：  每页条数 **约束限制**：  正整数 **取值范围**：  大于0，小于等于100 **默认取值**：  100
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：  分页下一页标志位 **约束限制**：  以列表查询接口返回为准 **取值范围**：  不涉及 **默认取值**：  不涉及
	Marker *string `json:"marker,omitempty"`
}

func (o ListDomainAllResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainAllResourceRequest struct{}"
	}

	return strings.Join([]string{"ListDomainAllResourceRequest", string(data)}, " ")
}
