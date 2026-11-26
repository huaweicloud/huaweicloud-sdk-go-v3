package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SpeicialConfiguration struct {

	// **参数解释：** 创建时间，Unix时间戳格式，单位为毫秒 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	CreateTime int64 `json:"create_time"`

	// **参数解释：** 备忘录 **约束限制：** 不涉及 **取值范围：** 不涉及
	Note string `json:"note"`
}

func (o SpeicialConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpeicialConfiguration struct{}"
	}

	return strings.Join([]string{"SpeicialConfiguration", string(data)}, " ")
}
