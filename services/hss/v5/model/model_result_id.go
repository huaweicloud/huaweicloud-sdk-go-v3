package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResultId **参数解释**： 病毒查杀结果ID **取值范围**： 字符长度1-64位
type ResultId struct {
}

func (o ResultId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultId struct{}"
	}

	return strings.Join([]string{"ResultId", string(data)}, " ")
}
