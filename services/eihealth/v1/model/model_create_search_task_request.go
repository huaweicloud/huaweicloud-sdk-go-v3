package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSearchTaskRequest Request Object
type CreateSearchTaskRequest struct {
	Body *SearchTaskData `json:"body,omitempty"`
}

func (o CreateSearchTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSearchTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateSearchTaskRequest", string(data)}, " ")
}
