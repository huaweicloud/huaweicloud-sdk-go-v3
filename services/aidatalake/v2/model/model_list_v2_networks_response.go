package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2NetworksResponse Response Object
type ListV2NetworksResponse struct {

	// **参数解释**：总记录数。 **取值范围**：不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**：网络配置列表。
	Networks       *[]V2NetworkBriefInfo `json:"networks,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListV2NetworksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2NetworksResponse struct{}"
	}

	return strings.Join([]string{"ListV2NetworksResponse", string(data)}, " ")
}
