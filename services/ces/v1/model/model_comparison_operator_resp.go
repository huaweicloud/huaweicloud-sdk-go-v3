package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComparisonOperatorResp **参数解释**： 阈值符号。     **取值范围**： 支持的值为(>|<|>=|<=|=|!=|cycle_decrease|cycle_increase|cycle_wave);cycle_decrease为环比下降,cycle_increase为环比上升,cycle_wave为环比波动。
type ComparisonOperatorResp struct {
}

func (o ComparisonOperatorResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComparisonOperatorResp struct{}"
	}

	return strings.Join([]string{"ComparisonOperatorResp", string(data)}, " ")
}
