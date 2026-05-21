package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRelatedDnsResponse Response Object
type ShowRelatedDnsResponse struct {

	// 关联DN。
	RelatedDataNodes *[]RelatedDnVo `json:"related_data_nodes,omitempty"`

	// 最近恢复时间点。
	LatestRestorableTime *string `json:"latest_restorable_time,omitempty"`

	// **参数解释**：  分页参数: 起始值。  **参数范围**：   大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：  分页参数: 每页记录数。  **参数范围**：  大于0且小于等于128。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：  总记录数。  **参数范围**：  不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowRelatedDnsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRelatedDnsResponse struct{}"
	}

	return strings.Join([]string{"ShowRelatedDnsResponse", string(data)}, " ")
}
