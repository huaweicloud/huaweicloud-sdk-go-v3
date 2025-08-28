package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BotRequestDistributionsNormalBucket **参数解释：** 正常请求的统计数据 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type BotRequestDistributionsNormalBucket struct {

	// **参数解释：** 正常请求总数 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 0
	Count *int64 `json:"count,omitempty"`
}

func (o BotRequestDistributionsNormalBucket) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BotRequestDistributionsNormalBucket struct{}"
	}

	return strings.Join([]string{"BotRequestDistributionsNormalBucket", string(data)}, " ")
}
