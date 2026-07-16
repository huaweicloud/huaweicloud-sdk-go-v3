package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerBindPublicIpRequest 绑定EIP请求体。
type ServerBindPublicIpRequest struct {

	// **参数解释**：EIP的ID。 **约束限制**：必填。 **取值范围**：1 - 64字符。 **默认取值**：不涉及。
	PublicIpId *string `json:"public_ip_id,omitempty"`
}

func (o ServerBindPublicIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerBindPublicIpRequest struct{}"
	}

	return strings.Join([]string{"ServerBindPublicIpRequest", string(data)}, " ")
}
