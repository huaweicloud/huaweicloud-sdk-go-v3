package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagResponse struct {

	// **参数解释：** 标签键。 **取值范围：** 不涉及。
	Key string `json:"key"`

	// **参数解释：** 标签值。 **取值范围：** 不涉及。
	Value string `json:"value"`
}

func (o TagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagResponse struct{}"
	}

	return strings.Join([]string{"TagResponse", string(data)}, " ")
}
