package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBinlogMergeResponse Response Object
type CreateBinlogMergeResponse struct {

	// **参数解释**：  合并Binlog任务ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateBinlogMergeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBinlogMergeResponse struct{}"
	}

	return strings.Join([]string{"CreateBinlogMergeResponse", string(data)}, " ")
}
