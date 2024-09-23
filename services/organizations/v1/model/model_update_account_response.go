package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccountResponse Response Object
type UpdateAccountResponse struct {
	Account        *AccountDto `json:"account,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o UpdateAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccountResponse struct{}"
	}

	return strings.Join([]string{"UpdateAccountResponse", string(data)}, " ")
}
