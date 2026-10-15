package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListRposAndRtosResponse Response Object
type BatchListRposAndRtosResponse struct {

	// **参数解释**：  返回的RPO和RTO查询结果总数，与results数组的长度一致。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Count *int32 `json:"count,omitempty"`

	// **参数解释**：  批量查询RPO和RTO的响应体集合，每个元素包含一个任务的RPO和RTO信息。详情请参见QueryRpoAndRtoResp对象定义。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Results        *[]QueryRpoAndRtoResp `json:"results,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o BatchListRposAndRtosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListRposAndRtosResponse struct{}"
	}

	return strings.Join([]string{"BatchListRposAndRtosResponse", string(data)}, " ")
}
