package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMergeRequestParticipantsResponse Response Object
type ListMergeRequestParticipantsResponse struct {
	Body           *[]UserBasicDto `json:"body,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListMergeRequestParticipantsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestParticipantsResponse struct{}"
	}

	return strings.Join([]string{"ListMergeRequestParticipantsResponse", string(data)}, " ")
}
