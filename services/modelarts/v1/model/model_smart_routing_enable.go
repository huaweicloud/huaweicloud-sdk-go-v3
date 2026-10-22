package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SmartRoutingEnable **参数解释：** 智能分流开启字段。 **取值范围：** [true, false]。
type SmartRoutingEnable struct {
}

func (o SmartRoutingEnable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartRoutingEnable struct{}"
	}

	return strings.Join([]string{"SmartRoutingEnable", string(data)}, " ")
}
