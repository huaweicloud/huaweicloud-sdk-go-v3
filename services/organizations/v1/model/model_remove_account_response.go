package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveAccountResponse Response Object
type RemoveAccountResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveAccountResponse struct{}"
	}

	return strings.Join([]string{"RemoveAccountResponse", string(data)}, " ")
}
