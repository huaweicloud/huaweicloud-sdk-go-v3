package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRemoveAccountsResponse Response Object
type BatchRemoveAccountsResponse struct {
	Data           *MultiAccountRespData `json:"data,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o BatchRemoveAccountsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemoveAccountsResponse struct{}"
	}

	return strings.Join([]string{"BatchRemoveAccountsResponse", string(data)}, " ")
}
