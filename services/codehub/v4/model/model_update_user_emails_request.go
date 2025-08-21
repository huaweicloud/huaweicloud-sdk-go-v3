package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserEmailsRequest Request Object
type UpdateUserEmailsRequest struct {
	Body *ModifyEmailAddressDto `json:"body,omitempty"`
}

func (o UpdateUserEmailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserEmailsRequest struct{}"
	}

	return strings.Join([]string{"UpdateUserEmailsRequest", string(data)}, " ")
}
