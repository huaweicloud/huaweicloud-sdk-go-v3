package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateApplicationViewResponse Response Object
type BatchCreateApplicationViewResponse struct {

	// **参数解释：** 应用id列表。 **取值范围：** 不涉及。
	ApplicationIds *[]string `json:"application_ids,omitempty"`

	// **参数解释：** 组件id列表。 **取值范围：** 不涉及。
	ComponentIds *[]string `json:"component_ids,omitempty"`

	// **参数解释：** 分组id列表。 **取值范围：** 不涉及。
	GroupIds       *[]string `json:"group_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchCreateApplicationViewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateApplicationViewResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateApplicationViewResponse", string(data)}, " ")
}
