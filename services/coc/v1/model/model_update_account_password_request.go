package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccountPasswordRequest Request Object
type UpdateAccountPasswordRequest struct {
	Body *UpdateAccountPasswordRequestBody `json:"body,omitempty"`
}

func (o UpdateAccountPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccountPasswordRequest struct{}"
	}

	return strings.Join([]string{"UpdateAccountPasswordRequest", string(data)}, " ")
}
