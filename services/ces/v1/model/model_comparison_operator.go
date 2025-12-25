package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComparisonOperator **参数解释**： 阈值符号。 **约束限制**： 指标告警可以使用的阈值符号有>、>=、<、<=、=、!=、cycle_decrease、cycle_increase、cycle_wave； 事件告警可以使用的阈值符号为>、>=、<、<=、=、!=。 **取值范围**： 支持的值为(>|<|>=|<=|=|!=|cycle_decrease|cycle_increase|cycle_wave);cycle_decrease为环比下降,cycle_increase为环比上升,cycle_wave为环比波动。
type ComparisonOperator struct {
}

func (o ComparisonOperator) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComparisonOperator struct{}"
	}

	return strings.Join([]string{"ComparisonOperator", string(data)}, " ")
}
