package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkloadsResponse Response Object
type ListWorkloadsResponse struct {

	// **参数解释**：资源的API版本。 **取值范围**：可选值如下： - v1：当前资源版本为v1
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**：资源的类型。 **取值范围**：可选值如下： - WorkloadList：作业列表
	Kind *string `json:"kind,omitempty"`

	// **参数解释**：资源池中的作业列表。
	Items          *[]Workload `json:"items,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListWorkloadsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkloadsResponse struct{}"
	}

	return strings.Join([]string{"ListWorkloadsResponse", string(data)}, " ")
}
