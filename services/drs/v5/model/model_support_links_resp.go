package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupportLinksResp 查询可用链路信息响应体
type SupportLinksResp struct {

	// **参数解释**：  列表中的项目总数，与分页无关。表示满足查询条件的记录总数量，用于分页计算。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**：  DRS支持的数据复制链路列表，每条链路包含引擎类型、网络类型、迁移模式等组合信息。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	SupportLinks *[]SupportLinkInfo `json:"support_links,omitempty"`
}

func (o SupportLinksResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupportLinksResp struct{}"
	}

	return strings.Join([]string{"SupportLinksResp", string(data)}, " ")
}
