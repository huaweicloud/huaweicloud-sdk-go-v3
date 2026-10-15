package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PageReq 分页请求体
type PageReq struct {

	// **参数解释**：  当前页码，用于分页查询时指定返回第几页的数据。填0可获取全部数据项。  **约束限制**：  - 当per_page填0时，cur_page只能填0或1。 - 不能超过传入jobID数量除每页任务数量的最大页。  **取值范围**：  不涉及。  **默认取值**：  1
	CurPage *int32 `json:"cur_page,omitempty"`

	// **参数解释**：  每页返回的记录数，用于分页查询时控制每页数据量。填0可获取全部数据项。  **约束限制**：  - 填0时，cur_page只能填0或1。  **取值范围**：  0-100。  **默认取值**：  5
	PerPage *int32 `json:"per_page,omitempty"`
}

func (o PageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageReq struct{}"
	}

	return strings.Join([]string{"PageReq", string(data)}, " ")
}
