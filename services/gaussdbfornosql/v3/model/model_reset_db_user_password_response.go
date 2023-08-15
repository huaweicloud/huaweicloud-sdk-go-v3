package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetDbUserPasswordResponse Response Object
type ResetDbUserPasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetDbUserPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetDbUserPasswordResponse struct{}"
	}

	return strings.Join([]string{"ResetDbUserPasswordResponse", string(data)}, " ")
}
