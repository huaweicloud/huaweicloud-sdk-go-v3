package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoDetect **参数解释**： 是否自动开启检测 **取值范围**: - true：是 - false：否
type AutoDetect struct {
}

func (o AutoDetect) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoDetect struct{}"
	}

	return strings.Join([]string{"AutoDetect", string(data)}, " ")
}
