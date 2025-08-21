package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryEventsResponse Response Object
type ListRepositoryEventsResponse struct {
	Body *[]RepositoryPushEventDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRepositoryEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryEventsResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryEventsResponse", string(data)}, " ")
}
