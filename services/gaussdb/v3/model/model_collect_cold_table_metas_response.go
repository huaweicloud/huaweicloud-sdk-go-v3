package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectColdTableMetasResponse Response Object
type CollectColdTableMetasResponse struct {

	// **参数解释**：  统计冷表元信息的任务ID。  **取值范围**：  不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CollectColdTableMetasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectColdTableMetasResponse struct{}"
	}

	return strings.Join([]string{"CollectColdTableMetasResponse", string(data)}, " ")
}
