package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaseResp 基础响应体。
type BaseResp struct {

	// **参数解释**：  列表中的项目总数，与分页无关。表示满足查询条件的记录总数量，用于分页计算。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	TotalCount int32 `json:"total_count"`
}

func (o BaseResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseResp struct{}"
	}

	return strings.Join([]string{"BaseResp", string(data)}, " ")
}
