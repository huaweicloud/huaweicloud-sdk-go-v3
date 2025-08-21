package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMergeRequestChangesResponse Response Object
type ListMergeRequestChangesResponse struct {
	Body           *[]SimpleMergeRequestChangesDto `json:"body,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListMergeRequestChangesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestChangesResponse struct{}"
	}

	return strings.Join([]string{"ListMergeRequestChangesResponse", string(data)}, " ")
}
