package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgencyServices struct {

	// **参数解释**：服务ID。服务ID，标识目标云服务（如 \"19\" 代表 EVS） **约束限制**：不涉及。 **取值范围**：长度限制为64个字符。 **默认取值**：空。
	ServiceId *string `json:"service_id,omitempty"`

	// **参数解释**：该服务下的细粒度权限操作列表，格式为 服务:资源:操作，如 \"evs:volumes:delete\" **约束限制**：不涉及。 **取值范围**：长度限制为256个字符。 **默认取值**：空。
	Actions *[]string `json:"actions,omitempty"`
}

func (o AgencyServices) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyServices struct{}"
	}

	return strings.Join([]string{"AgencyServices", string(data)}, " ")
}
