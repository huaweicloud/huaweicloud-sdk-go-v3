package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DomainsStatusesRequestBody DomainsStatuses请求体
type DomainsStatusesRequestBody struct {

	// **参数解释**： 租户id集合。 **约束限制**： 不涉及。 **取值范围**： 不涉及。
	DomainIds *string `json:"domain_ids,omitempty"`
}

func (o DomainsStatusesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainsStatusesRequestBody struct{}"
	}

	return strings.Join([]string{"DomainsStatusesRequestBody", string(data)}, " ")
}
