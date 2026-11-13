package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishApplicationRequest Request Object
type PublishApplicationRequest struct {

	// **参数解释**： 租户ID。 **约束限制**： 不涉及。 **取值范围**： [1-50]个字符。 **默认取值**： 不涉及。
	TenantId string `json:"tenant_id"`

	// **参数解释**： 应用ID。 **约束限制**： 不涉及。 **取值范围**： [1-256]个字符。 **默认取值**： 不涉及。
	ApplicationId string `json:"application_id"`

	Body *ReleaseAppReq `json:"body,omitempty"`
}

func (o PublishApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishApplicationRequest struct{}"
	}

	return strings.Join([]string{"PublishApplicationRequest", string(data)}, " ")
}
