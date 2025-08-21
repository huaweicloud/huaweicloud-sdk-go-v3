package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InspectionVpcInfo struct {

	// **参数解释**： 引流VPC  **取值范围**： 不涉及
	InspectionVpcId *string `json:"inspection_vpc_id,omitempty"`

	// **参数解释**： 引流VPC  **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`
}

func (o InspectionVpcInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InspectionVpcInfo struct{}"
	}

	return strings.Join([]string{"InspectionVpcInfo", string(data)}, " ")
}
