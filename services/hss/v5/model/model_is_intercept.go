package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsIntercept **参数解释**： 是否开启阻断 **取值范围**: - true：是 - false：否
type IsIntercept struct {
}

func (o IsIntercept) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsIntercept struct{}"
	}

	return strings.Join([]string{"IsIntercept", string(data)}, " ")
}
