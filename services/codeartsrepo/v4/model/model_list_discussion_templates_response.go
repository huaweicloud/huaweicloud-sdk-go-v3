package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDiscussionTemplatesResponse Response Object
type ListDiscussionTemplatesResponse struct {
	Body           *[]DiscussionTemplateDto `json:"body,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListDiscussionTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDiscussionTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListDiscussionTemplatesResponse", string(data)}, " ")
}
