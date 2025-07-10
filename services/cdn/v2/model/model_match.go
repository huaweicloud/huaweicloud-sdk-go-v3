package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Match **参数解释：** 规则匹配条件 **约束限制：** 不涉及
type Match struct {

	// **参数解释：** 逻辑运算符 **约束限制：** 不涉及 **取值范围：** - and: 与关系 - or: 或关系 **默认取值：** 不涉及
	Logic string `json:"logic"`

	// **参数解释：** 匹配条件列表 **约束限制：** 不涉及
	Criteria []Criteria `json:"criteria"`
}

func (o Match) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Match struct{}"
	}

	return strings.Join([]string{"Match", string(data)}, " ")
}
