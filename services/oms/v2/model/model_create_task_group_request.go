package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTaskGroupRequest Request Object
type CreateTaskGroupRequest struct {
	Body *CreateTaskGroupReq `json:"body,omitempty"`
}

func (o CreateTaskGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTaskGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateTaskGroupRequest", string(data)}, " ")
}
