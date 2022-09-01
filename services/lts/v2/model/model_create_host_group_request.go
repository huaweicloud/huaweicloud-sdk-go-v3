package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateHostGroupRequest struct {
	Body *CreateHostGroupRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateHostGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateHostGroupRequest", string(data)}, " ")
}
