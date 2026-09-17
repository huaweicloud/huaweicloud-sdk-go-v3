package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUpgradeWorkFlowsResponse Response Object
type ListUpgradeWorkFlowsResponse struct {

	// **参数解释：** API类型，固定值\"List\"，该值不可修改 **约束限制：** 固定值 **取值范围：** - List  **默认取值：** List
	Kind *string `json:"kind,omitempty"`

	// **参数解释：** API版本，固定值\"v3\"，该值不可修改 **约束限制：** 固定值 **取值范围：** - v3  **默认取值：** v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释：** 升级工作流列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Items          *[]UpgradeWorkFlow `json:"items,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListUpgradeWorkFlowsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUpgradeWorkFlowsResponse struct{}"
	}

	return strings.Join([]string{"ListUpgradeWorkFlowsResponse", string(data)}, " ")
}
