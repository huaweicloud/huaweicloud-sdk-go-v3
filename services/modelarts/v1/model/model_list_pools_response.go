package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPoolsResponse Response Object
type ListPoolsResponse struct {

	// **参数解释**：资源的API版本。 **取值范围**：可选值如下： - v2：当前资源版本为v2。
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**：资源的类型。 **取值范围**：可选值如下： - PoolList：资源池列表。
	Kind *string `json:"kind,omitempty"`

	// **参数解释**：资源池列表。
	Items          *[]PoolModel `json:"items,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListPoolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoolsResponse struct{}"
	}

	return strings.Join([]string{"ListPoolsResponse", string(data)}, " ")
}
