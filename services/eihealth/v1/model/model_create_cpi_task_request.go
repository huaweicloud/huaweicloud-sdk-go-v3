package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCpiTaskRequest Request Object
type CreateCpiTaskRequest struct {
	Body *CpiTaskData `json:"body,omitempty"`
}

func (o CreateCpiTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCpiTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateCpiTaskRequest", string(data)}, " ")
}
