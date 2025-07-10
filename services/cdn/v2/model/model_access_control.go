package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessControl **参数解释：** 访问控制 **约束限制：** 不涉及
type AccessControl struct {

	// **参数解释：** 访问控制类型 **约束限制：** 不涉及 **取值范围：** - block: 拒绝 - trust: 允许 **默认取值：** 不涉及
	Type string `json:"type"`
}

func (o AccessControl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessControl struct{}"
	}

	return strings.Join([]string{"AccessControl", string(data)}, " ")
}
