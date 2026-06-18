package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectMergeRequestTemplatesResponse Response Object
type ListProjectMergeRequestTemplatesResponse struct {
	Body *[]ProjectMergeRequestTemplateDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListProjectMergeRequestTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectMergeRequestTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListProjectMergeRequestTemplatesResponse", string(data)}, " ")
}
