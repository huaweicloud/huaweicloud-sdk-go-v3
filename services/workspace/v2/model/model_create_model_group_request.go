package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateModelGroupRequest Request Object
type CreateModelGroupRequest struct {
	Body *CreateModelGroupReq `json:"body,omitempty"`
}

func (o CreateModelGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateModelGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateModelGroupRequest", string(data)}, " ")
}
