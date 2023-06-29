package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeShareNameResponse Response Object
type ChangeShareNameResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeShareNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeShareNameResponse struct{}"
	}

	return strings.Join([]string{"ChangeShareNameResponse", string(data)}, " ")
}
