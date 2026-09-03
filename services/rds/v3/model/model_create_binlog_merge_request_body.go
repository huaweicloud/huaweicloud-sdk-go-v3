package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBinlogMergeRequestBody 合并Binlog请求体
type CreateBinlogMergeRequestBody struct {

	// **参数解释**：  查询开始时间，格式为Unix时间戳，单位为毫秒。  **约束限制**：  开始时间需早于结束时间。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	StartTime int64 `json:"start_time"`

	// **参数解释**：  查询结束时间，格式为Unix时间戳，单位为毫秒。  **约束限制**：  结束时间需晚于开始时间。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	EndTime int64 `json:"end_time"`
}

func (o CreateBinlogMergeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBinlogMergeRequestBody struct{}"
	}

	return strings.Join([]string{"CreateBinlogMergeRequestBody", string(data)}, " ")
}
