package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2ComputesResponse Response Object
type ListV2ComputesResponse struct {

	// **参数解释**：总记录数。 **取值范围**：0~10000。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**：计算资源列表，每个元素为一个V2ComputeInfo对象，包含计算资源ID、名称、计费模式、容量配置、标签等详细信息。
	Computes       *[]V2ComputeInfo `json:"computes,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListV2ComputesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2ComputesResponse struct{}"
	}

	return strings.Join([]string{"ListV2ComputesResponse", string(data)}, " ")
}
