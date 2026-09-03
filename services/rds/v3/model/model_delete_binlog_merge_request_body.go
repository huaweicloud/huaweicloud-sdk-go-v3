package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBinlogMergeRequestBody 删除Binlog合并记录请求体
type DeleteBinlogMergeRequestBody struct {

	// **参数解释**：  Binlog合并记录ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Id string `json:"id"`
}

func (o DeleteBinlogMergeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBinlogMergeRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteBinlogMergeRequestBody", string(data)}, " ")
}
