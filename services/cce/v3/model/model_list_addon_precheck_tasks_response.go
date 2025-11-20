package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAddonPrecheckTasksResponse Response Object
type ListAddonPrecheckTasksResponse struct {

	// **参数解释：** API类型 **取值范围：** 固定值\"AddonCheck\"
	Kind *string `json:"kind,omitempty"`

	// **参数解释：** API版本 **取值范围：** 固定值\"v3\"
	ApiVersion *string `json:"apiVersion,omitempty"`

	PageInfo *PageInfo `json:"pageInfo,omitempty"`

	// **参数解释：** 插件检查任务信息列表，包含了插件检查任务ID，插件模板名称，插件实例ID等。 **约束限制：** 不涉及
	Items          *[]AddonCheckTask `json:"items,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListAddonPrecheckTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddonPrecheckTasksResponse struct{}"
	}

	return strings.Join([]string{"ListAddonPrecheckTasksResponse", string(data)}, " ")
}
