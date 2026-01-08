package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResPortRange 端口监听范围（闭区间)，最多指定20个端口组，每个组范围不可有重叠部分 仅当protocol_port为0时可以传入。
type ResPortRange struct {

	// **参数解释**：起始端口。 **取值范围**：1-65535
	StartPort int32 `json:"start_port"`

	// **参数解释**：结束端口。 **取值范围**：1-65535
	EndPort int32 `json:"end_port"`
}

func (o ResPortRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResPortRange struct{}"
	}

	return strings.Join([]string{"ResPortRange", string(data)}, " ")
}
