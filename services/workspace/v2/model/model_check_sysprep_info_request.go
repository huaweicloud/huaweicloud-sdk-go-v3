package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckSysprepInfoRequest Request Object
type CheckSysprepInfoRequest struct {
	Body *CheckSysprepInfoRequestBody `json:"body,omitempty"`
}

func (o CheckSysprepInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckSysprepInfoRequest struct{}"
	}

	return strings.Join([]string{"CheckSysprepInfoRequest", string(data)}, " ")
}
