package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceTablesForHtapResponse Response Object
type ShowInstanceTablesForHtapResponse struct {

	// **参数解释**：  数据库表名称。  **默认取值**：  不涉及。
	Tables *[]string `json:"tables,omitempty"`

	// **参数解释**：  数据库数量。  **默认取值**：  不涉及。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowInstanceTablesForHtapResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceTablesForHtapResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceTablesForHtapResponse", string(data)}, " ")
}
