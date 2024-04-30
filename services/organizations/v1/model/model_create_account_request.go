package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAccountRequest Request Object
type CreateAccountRequest struct {
	Body *CreateAccountReqBody `json:"body,omitempty"`
}

func (o CreateAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccountRequest struct{}"
	}

	return strings.Join([]string{"CreateAccountRequest", string(data)}, " ")
}
