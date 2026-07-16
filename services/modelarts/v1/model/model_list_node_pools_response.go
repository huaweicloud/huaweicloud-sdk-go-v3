package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNodePoolsResponse Response Object
type ListNodePoolsResponse struct {

	// **参数解释**： API版本。 **取值范围**： 可选值如下： - v2
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**： 资源类型。 **取值范围**： 可选值如下： - NodePoolList：节点列表
	Kind *string `json:"kind,omitempty"`

	// **参数解释**：节点池列表。
	Items          *[]NodePool `json:"items,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListNodePoolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodePoolsResponse struct{}"
	}

	return strings.Join([]string{"ListNodePoolsResponse", string(data)}, " ")
}
