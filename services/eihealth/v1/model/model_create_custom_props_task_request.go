package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCustomPropsTaskRequest Request Object
type CreateCustomPropsTaskRequest struct {
	Body *CustomPropsTaskData `json:"body,omitempty"`
}

func (o CreateCustomPropsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomPropsTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateCustomPropsTaskRequest", string(data)}, " ")
}
