package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindDevServerPublicIpRequest Request Object
type BindDevServerPublicIpRequest struct {

	// **参数解释**：Lite Server ID。 **约束限制**：必填。 **取值范围**：1 - 64字符。 **默认取值**：不涉及。
	Id string `json:"id"`

	Body *ServerBindPublicIpRequest `json:"body,omitempty"`
}

func (o BindDevServerPublicIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindDevServerPublicIpRequest struct{}"
	}

	return strings.Join([]string{"BindDevServerPublicIpRequest", string(data)}, " ")
}
