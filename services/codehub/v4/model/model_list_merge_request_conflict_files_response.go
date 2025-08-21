package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMergeRequestConflictFilesResponse Response Object
type ListMergeRequestConflictFilesResponse struct {
	Body           *[]MrConflictFileDto `json:"body,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListMergeRequestConflictFilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestConflictFilesResponse struct{}"
	}

	return strings.Join([]string{"ListMergeRequestConflictFilesResponse", string(data)}, " ")
}
