package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlinkSavepointResponse Response Object
type ListFlinkSavepointResponse struct {

	// **参数解释**：Savepoint总数。 **取值范围**：0~2147483647。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**：Savepoint摘要信息列表。可显示0~1000个FlinkSavepointSummary对象。
	Savepoints     *[]FlinkSavepointSummary `json:"savepoints,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListFlinkSavepointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlinkSavepointResponse struct{}"
	}

	return strings.Join([]string{"ListFlinkSavepointResponse", string(data)}, " ")
}
