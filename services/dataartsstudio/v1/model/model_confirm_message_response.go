package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmMessageResponse Response Object
type ConfirmMessageResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ConfirmMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmMessageResponse struct{}"
	}

	return strings.Join([]string{"ConfirmMessageResponse", string(data)}, " ")
}
