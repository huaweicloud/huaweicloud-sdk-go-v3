package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompleteResourceTokenAuthResponse Response Object
type CompleteResourceTokenAuthResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CompleteResourceTokenAuthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompleteResourceTokenAuthResponse struct{}"
	}

	return strings.Join([]string{"CompleteResourceTokenAuthResponse", string(data)}, " ")
}
