package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAccountResponse struct {
	Account        *AccountDto `json:"account,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccountResponse struct{}"
	}

	return strings.Join([]string{"ShowAccountResponse", string(data)}, " ")
}
