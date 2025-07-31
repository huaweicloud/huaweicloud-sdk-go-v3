package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetAccountPasswordRequest Request Object
type ResetAccountPasswordRequest struct {
	Body *ResetAccountPasswordRequestBody `json:"body,omitempty"`
}

func (o ResetAccountPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetAccountPasswordRequest struct{}"
	}

	return strings.Join([]string{"ResetAccountPasswordRequest", string(data)}, " ")
}
