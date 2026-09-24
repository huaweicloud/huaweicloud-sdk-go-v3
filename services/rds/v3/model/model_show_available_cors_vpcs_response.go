package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailableCorsVpcsResponse Response Object
type ShowAvailableCorsVpcsResponse struct {

	// **参数解释**：  VPC ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**：  子网ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	SubnetId *string `json:"subnet_id,omitempty"`

	// **参数解释**：  安全组ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	SecurityGroupId *string `json:"security_group_id,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o ShowAvailableCorsVpcsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableCorsVpcsResponse struct{}"
	}

	return strings.Join([]string{"ShowAvailableCorsVpcsResponse", string(data)}, " ")
}
