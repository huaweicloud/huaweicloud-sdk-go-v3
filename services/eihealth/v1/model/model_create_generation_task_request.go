package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGenerationTaskRequest Request Object
type CreateGenerationTaskRequest struct {
	Body *GenerationTaskData `json:"body,omitempty"`
}

func (o CreateGenerationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGenerationTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateGenerationTaskRequest", string(data)}, " ")
}
