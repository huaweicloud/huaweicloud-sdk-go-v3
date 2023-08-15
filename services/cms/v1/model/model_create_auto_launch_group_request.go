package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAutoLaunchGroupRequest Request Object
type CreateAutoLaunchGroupRequest struct {
	Body *AutoLaunchGroupReqV2 `json:"body,omitempty"`
}

func (o CreateAutoLaunchGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAutoLaunchGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateAutoLaunchGroupRequest", string(data)}, " ")
}
