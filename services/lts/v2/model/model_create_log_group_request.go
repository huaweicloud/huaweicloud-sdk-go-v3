package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateLogGroupRequest struct {
	Body *CreateLogGroupParams `json:"body,omitempty"`
}

func (o CreateLogGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateLogGroupRequest", string(data)}, " ")
}
