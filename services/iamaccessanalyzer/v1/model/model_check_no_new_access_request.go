package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckNoNewAccessRequest Request Object
type CheckNoNewAccessRequest struct {
	Body *CheckNoNewAccessReqBody `json:"body,omitempty"`
}

func (o CheckNoNewAccessRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckNoNewAccessRequest struct{}"
	}

	return strings.Join([]string{"CheckNoNewAccessRequest", string(data)}, " ")
}
