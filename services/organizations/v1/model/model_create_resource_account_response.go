package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceAccountResponse Response Object
type CreateResourceAccountResponse struct {
	CreateAccountStatus *CreateAccountStatusDto `json:"create_account_status,omitempty"`
	HttpStatusCode      int                     `json:"-"`
}

func (o CreateResourceAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceAccountResponse struct{}"
	}

	return strings.Join([]string{"CreateResourceAccountResponse", string(data)}, " ")
}
