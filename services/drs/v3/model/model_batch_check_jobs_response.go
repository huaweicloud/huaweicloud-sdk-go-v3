package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCheckJobsResponse Response Object
type BatchCheckJobsResponse struct {

	// **参数解释**：  预检查响应体，详情请参见PostPreCheckResp对象定义。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Results *[]PostPreCheckResp `json:"results,omitempty"`

	// **参数解释**：  返回的预检查结果总数，与results数组的长度一致。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchCheckJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchCheckJobsResponse", string(data)}, " ")
}
