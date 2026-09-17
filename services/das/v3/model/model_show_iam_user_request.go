package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIamUserRequest Request Object
type ShowIamUserRequest struct {
	Body *ShowIamUserRequestBody `json:"body,omitempty"`
}

func (o ShowIamUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIamUserRequest struct{}"
	}

	return strings.Join([]string{"ShowIamUserRequest", string(data)}, " ")
}
