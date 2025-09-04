package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckAppGroupRequest Request Object
type CheckAppGroupRequest struct {
	Body *CheckAppGroupReq `json:"body,omitempty"`
}

func (o CheckAppGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckAppGroupRequest struct{}"
	}

	return strings.Join([]string{"CheckAppGroupRequest", string(data)}, " ")
}
