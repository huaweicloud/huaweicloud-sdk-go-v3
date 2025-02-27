package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserPasswordRequest Request Object
type UpdateUserPasswordRequest struct {
	Body *ChangeUsersPassword `json:"body,omitempty"`
}

func (o UpdateUserPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserPasswordRequest struct{}"
	}

	return strings.Join([]string{"UpdateUserPasswordRequest", string(data)}, " ")
}
