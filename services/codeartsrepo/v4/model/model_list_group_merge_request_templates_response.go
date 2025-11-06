package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupMergeRequestTemplatesResponse Response Object
type ListGroupMergeRequestTemplatesResponse struct {
	Body           *[]GroupMergeRequestTemplateDto `json:"body,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListGroupMergeRequestTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupMergeRequestTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListGroupMergeRequestTemplatesResponse", string(data)}, " ")
}
