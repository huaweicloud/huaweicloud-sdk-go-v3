package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateObjectProcessTaskRequest Request Object
type CreateObjectProcessTaskRequest struct {
	Body *ObjectProcessReq `json:"body,omitempty"`
}

func (o CreateObjectProcessTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateObjectProcessTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateObjectProcessTaskRequest", string(data)}, " ")
}
