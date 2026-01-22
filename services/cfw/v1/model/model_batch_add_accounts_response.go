package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddAccountsResponse Response Object
type BatchAddAccountsResponse struct {
	Data           *MultiAccountRespData `json:"data,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o BatchAddAccountsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddAccountsResponse struct{}"
	}

	return strings.Join([]string{"BatchAddAccountsResponse", string(data)}, " ")
}
