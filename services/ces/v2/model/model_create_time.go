package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTime **参数解释**： 告警模板的创建时间 **取值范围**： 不涉及。
type CreateTime struct {
}

func (o CreateTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTime struct{}"
	}

	return strings.Join([]string{"CreateTime", string(data)}, " ")
}
