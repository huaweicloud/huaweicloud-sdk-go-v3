package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StoryPoint struct {

	// **参数解释：** 故事点id。 **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 故事点名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o StoryPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StoryPoint struct{}"
	}

	return strings.Join([]string{"StoryPoint", string(data)}, " ")
}
