package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAccountResponse Response Object
type CreateAccountResponse struct {
	CreateAccountStatus *CreateAccountStatusDto `json:"create_account_status,omitempty"`
	HttpStatusCode      int                     `json:"-"`
}

func (o CreateAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccountResponse struct{}"
	}

	return strings.Join([]string{"CreateAccountResponse", string(data)}, " ")
}
