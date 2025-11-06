package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComparisonOperatorResp **参数解释**： 告警阈值的比较条件。 **取值范围**： 只能是>、=、<、>=、<=、!=。
type ComparisonOperatorResp struct {
}

func (o ComparisonOperatorResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComparisonOperatorResp struct{}"
	}

	return strings.Join([]string{"ComparisonOperatorResp", string(data)}, " ")
}
