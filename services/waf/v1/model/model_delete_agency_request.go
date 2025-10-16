package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAgencyRequest Request Object
type DeleteAgencyRequest struct {

	// **参数解释：** 独享引擎操作 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Purged *bool `json:"purged,omitempty"`

	// **参数解释：** 待删除的代理id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	RoleIdList *[]string `json:"role_id_list,omitempty"`
}

func (o DeleteAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAgencyRequest struct{}"
	}

	return strings.Join([]string{"DeleteAgencyRequest", string(data)}, " ")
}
