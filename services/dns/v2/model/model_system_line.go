package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SystemLine struct {

	// **参数解释：** 线路ID。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 父线路ID。 **取值范围：** 不涉及。
	FatherId *string `json:"father_id,omitempty"`

	// **参数解释：** 线路名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o SystemLine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SystemLine struct{}"
	}

	return strings.Join([]string{"SystemLine", string(data)}, " ")
}
