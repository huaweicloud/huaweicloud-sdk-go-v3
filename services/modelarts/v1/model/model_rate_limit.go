package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RateLimit **参数解释：** 流量控制配置。 **约束限制：** 不涉及。
type RateLimit struct {

	// **参数解释：** 服务流量限制是指指定单位内一个服务能够被访问的次数上限。 **约束限制：** 不涉及。 **取值范围：** 1-10000。 **默认取值：** 不涉及。
	Num int32 `json:"num"`

	// **参数解释：** 流量限制单元。 **约束限制：** 不涉及。 **取值范围：** - SECONDS（秒）。 - MINUTES（分钟）。 **默认取值：** 不涉及。
	Unit string `json:"unit"`
}

func (o RateLimit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RateLimit struct{}"
	}

	return strings.Join([]string{"RateLimit", string(data)}, " ")
}
