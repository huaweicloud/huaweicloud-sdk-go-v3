package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBlocksRequest Request Object
type ShowBlocksRequest struct {

	// **参数解释**： 租户ID。 **约束限制**： 不涉及。 **取值范围**： [1-50]个字符。 **默认取值**： 不涉及。
	TenantId string `json:"tenant_id"`
}

func (o ShowBlocksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBlocksRequest struct{}"
	}

	return strings.Join([]string{"ShowBlocksRequest", string(data)}, " ")
}
