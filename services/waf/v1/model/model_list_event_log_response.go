package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventLogResponse Response Object
type ListEventLogResponse struct {
	Total *int64 `json:"total,omitempty"`

	Items          *[]EventDump `json:"items,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListEventLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventLogResponse struct{}"
	}

	return strings.Join([]string{"ListEventLogResponse", string(data)}, " ")
}
