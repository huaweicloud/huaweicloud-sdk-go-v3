package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationDetailsRequest Request Object
type ListApplicationDetailsRequest struct {

	// **参数解释**： 应用ID。 **约束限制**： 不涉及。 **取值范围**： [1-256]个字符。 **默认取值**： 不涉及。
	ApplicationId string `json:"application_id"`

	// **参数解释**： 查询类型。 **约束限制**： 不涉及。 **取值范围**： - 0：应用基本信息。 - 1：系统配置信息。 - 2：文档块信息。 - 3：发布信息。 - 4：所有信息 **默认取值**： 0。
	QueryType *string `json:"query_type,omitempty"`

	// **参数解释**： 页码。 **约束限制**： 仅当query_type为3或者4时该字段有效。 **取值范围**： 不涉及。 **默认取值**： 1。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 单页数量。 **约束限制**： 仅当query_type为3或者4时该字段有效。 **取值范围**： [1-100]。 **默认取值**： 20。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 租户ID。 **约束限制**： 不涉及。 **取值范围**： [1-50]个字符。 **默认取值**： 不涉及。
	TenantId string `json:"tenant_id"`
}

func (o ListApplicationDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListApplicationDetailsRequest", string(data)}, " ")
}
