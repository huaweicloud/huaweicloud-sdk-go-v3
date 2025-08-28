package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TypedStatBucket struct {

	// **参数解释：** 统计键（如bot类型、域名等） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释：** 该键对应的命中请求数量 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 0
	Count *int64 `json:"count,omitempty"`
}

func (o TypedStatBucket) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TypedStatBucket struct{}"
	}

	return strings.Join([]string{"TypedStatBucket", string(data)}, " ")
}
