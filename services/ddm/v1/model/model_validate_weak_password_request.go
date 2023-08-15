package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateWeakPasswordRequest Request Object
type ValidateWeakPasswordRequest struct {
	Body *WeakPasswordReq `json:"body,omitempty"`
}

func (o ValidateWeakPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateWeakPasswordRequest struct{}"
	}

	return strings.Join([]string{"ValidateWeakPasswordRequest", string(data)}, " ")
}
