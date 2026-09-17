package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatusAttributeVo struct {

	// **参数解释：** 项目id。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o StatusAttributeVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatusAttributeVo struct{}"
	}

	return strings.Join([]string{"StatusAttributeVo", string(data)}, " ")
}
