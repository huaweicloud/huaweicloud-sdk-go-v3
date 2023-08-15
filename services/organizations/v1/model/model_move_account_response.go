package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MoveAccountResponse Response Object
type MoveAccountResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o MoveAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoveAccountResponse struct{}"
	}

	return strings.Join([]string{"MoveAccountResponse", string(data)}, " ")
}
