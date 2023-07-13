package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppGroupRequest Request Object
type CreateAppGroupRequest struct {
	Body *CreateAppGroupReq `json:"body,omitempty"`
}

func (o CreateAppGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateAppGroupRequest", string(data)}, " ")
}
