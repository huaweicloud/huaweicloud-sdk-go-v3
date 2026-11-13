package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApplicationRequest Request Object
type CreateApplicationRequest struct {

	// **参数解释**： 用户租户ID。 **约束限制**： 不涉及。 **取值范围**： [1-50]个字符。 **默认取值**： 不涉及。
	TenantId string `json:"tenant_id"`

	Body *CreateAppReq `json:"body,omitempty"`
}

func (o CreateApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationRequest struct{}"
	}

	return strings.Join([]string{"CreateApplicationRequest", string(data)}, " ")
}
