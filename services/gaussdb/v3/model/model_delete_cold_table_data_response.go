package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteColdTableDataResponse Response Object
type DeleteColdTableDataResponse struct {

	// **参数解释**：  删除冷表数据的任务ID。  **取值范围**：  不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteColdTableDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteColdTableDataResponse struct{}"
	}

	return strings.Join([]string{"DeleteColdTableDataResponse", string(data)}, " ")
}
