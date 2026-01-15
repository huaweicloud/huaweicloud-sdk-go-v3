package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckCidrRequest Request Object
type CheckCidrRequest struct {
	Body *CheckCidrRequestBody `json:"body,omitempty"`
}

func (o CheckCidrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckCidrRequest struct{}"
	}

	return strings.Join([]string{"CheckCidrRequest", string(data)}, " ")
}
