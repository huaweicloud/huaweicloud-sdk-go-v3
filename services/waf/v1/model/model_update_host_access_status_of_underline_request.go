package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHostAccessStatusOfUnderlineRequest Request Object
type UpdateHostAccessStatusOfUnderlineRequest struct {

	// **参数解释：** 域名Id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	InstanceId string `json:"instance_id"`

	Body *AccessStatusBody `json:"body,omitempty"`
}

func (o UpdateHostAccessStatusOfUnderlineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostAccessStatusOfUnderlineRequest struct{}"
	}

	return strings.Join([]string{"UpdateHostAccessStatusOfUnderlineRequest", string(data)}, " ")
}
