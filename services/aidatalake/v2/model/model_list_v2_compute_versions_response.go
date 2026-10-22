package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2ComputeVersionsResponse Response Object
type ListV2ComputeVersionsResponse struct {

	// **参数解释**：总记录数。 **取值范围**：0~10000。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**：计算资源操作历史列表，列表元素为计算资源操作历史详情对象。
	ComputeVersions *[]ShowV2ComputeVersionResponseBody `json:"compute_versions,omitempty"`
	HttpStatusCode  int                                 `json:"-"`
}

func (o ListV2ComputeVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2ComputeVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListV2ComputeVersionsResponse", string(data)}, " ")
}
