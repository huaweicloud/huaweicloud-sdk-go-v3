/*
 * DMS
 *
 * DMS Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListConsumerGroupsRequest struct {
	ProjectId          string `json:"project_id"`
	QueueId            string `json:"queue_id"`
	IncludeDeadletter  *bool  `json:"include_deadletter,omitempty"`
	IncludeMessagesNum *bool  `json:"include_messages_num,omitempty"`
	PageSize           *int32 `json:"page_size,omitempty"`
	CurrentPage        *int32 `json:"current_page,omitempty"`
}

func (o ListConsumerGroupsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListConsumerGroupsRequest", string(data)}, " ")
}
