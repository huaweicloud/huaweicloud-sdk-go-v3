package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMergeRequestTemplatesResponse Response Object
type ListMergeRequestTemplatesResponse struct {
	Body           *[]MergeRequestTemplateDto `json:"body,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListMergeRequestTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListMergeRequestTemplatesResponse", string(data)}, " ")
}
