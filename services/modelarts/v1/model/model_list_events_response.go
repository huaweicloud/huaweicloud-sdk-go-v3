package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventsResponse Response Object
type ListEventsResponse struct {

	// **参数描述**：API版本。 **取值范围**：可选值如下： - v1
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数描述**：资源类型。 **取值范围**：可选值如下： - EventList：事件列表
	Kind *string `json:"kind,omitempty"`

	Metadata *EventListMeta `json:"metadata,omitempty"`

	// **参数描述**：事件列表。
	Items          *[]EventModel `json:"items,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventsResponse struct{}"
	}

	return strings.Join([]string{"ListEventsResponse", string(data)}, " ")
}
