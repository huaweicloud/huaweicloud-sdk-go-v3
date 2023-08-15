package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeUserStatusResponse Response Object
type ChangeUserStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeUserStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeUserStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangeUserStatusResponse", string(data)}, " ")
}
