package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuotasRequest Request Object
type ShowQuotasRequest struct {

	// **参数解释**： 是否包含标签配额标志。 **约束限制**： 不涉及。 **取值范围**： - true：包含配额。 - false：不包含配额。 **默认取值**： true。
	IncludeTagsQuota *string `json:"includeTagsQuota,omitempty"`

	// **参数解释**： 查询指定配额引擎。 **约束限制**： 不涉及。 **取值范围**： - reliability：RocketMQ消息引擎别称。 - kafka：kafka消息引擎别称。 **默认取值**： 不涉及。
	OnlyQuota *string `json:"onlyQuota,omitempty"`
}

func (o ShowQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowQuotasRequest", string(data)}, " ")
}
