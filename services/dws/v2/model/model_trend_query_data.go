package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TrendQueryData struct {

	// **参数解释**： 查询结果。 **取值范围**： 不涉及。
	Result *string `json:"result,omitempty"`

	// **参数解释**： 时间戳。 **取值范围**： 不涉及。
	Timestamp *int64 `json:"timestamp,omitempty"`
}

func (o TrendQueryData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrendQueryData struct{}"
	}

	return strings.Join([]string{"TrendQueryData", string(data)}, " ")
}
