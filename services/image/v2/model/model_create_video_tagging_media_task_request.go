package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVideoTaggingMediaTaskRequest Request Object
type CreateVideoTaggingMediaTaskRequest struct {
	Body *CreateVideoTaggingMediaTaskRequestBody `json:"body,omitempty"`
}

func (o CreateVideoTaggingMediaTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoTaggingMediaTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateVideoTaggingMediaTaskRequest", string(data)}, " ")
}
