package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateWorkSpaceRequest struct {
	Body *RequestWorkspace `json:"body,omitempty"`
}

func (o CreateWorkSpaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkSpaceRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkSpaceRequest", string(data)}, " ")
}
