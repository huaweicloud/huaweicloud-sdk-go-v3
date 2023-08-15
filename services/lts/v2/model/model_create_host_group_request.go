package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHostGroupRequest Request Object
type CreateHostGroupRequest struct {
	Body *CreateHostGroupRequestBody `json:"body,omitempty"`
}

func (o CreateHostGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateHostGroupRequest", string(data)}, " ")
}
