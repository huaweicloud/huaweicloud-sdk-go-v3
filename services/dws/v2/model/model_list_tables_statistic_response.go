package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTablesStatisticResponse Response Object
type ListTablesStatisticResponse struct {

	// **参数解释**： 数据采集时间毫秒级时间戳。 **取值范围**： 不涉及。
	CollectTime *int64 `json:"collect_time,omitempty"`

	// **参数解释**： 表倾斜率或脏页率列表。 **取值范围**： 不涉及。
	Data *[]ListTablesStatisticDto `json:"data,omitempty"`

	// **参数解释**： 总列表大小。 **取值范围**： 不涉及。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTablesStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTablesStatisticResponse struct{}"
	}

	return strings.Join([]string{"ListTablesStatisticResponse", string(data)}, " ")
}
