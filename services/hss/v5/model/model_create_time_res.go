package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTimeRes **参数解释**： 创建时间 **取值范围**： 最小值0，最大值9223372036854775807
type CreateTimeRes struct {
}

func (o CreateTimeRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTimeRes struct{}"
	}

	return strings.Join([]string{"CreateTimeRes", string(data)}, " ")
}
