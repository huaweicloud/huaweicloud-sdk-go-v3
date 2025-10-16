package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePremiumInstanceProgressRequest Request Object
type UpdatePremiumInstanceProgressRequest struct {

	// **参数解释：** 域名Id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	HostId string `json:"host_id"`

	Body *AccessProgress `json:"body,omitempty"`
}

func (o UpdatePremiumInstanceProgressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePremiumInstanceProgressRequest struct{}"
	}

	return strings.Join([]string{"UpdatePremiumInstanceProgressRequest", string(data)}, " ")
}
