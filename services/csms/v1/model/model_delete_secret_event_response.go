package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecretEventResponse Response Object
type DeleteSecretEventResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecretEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecretEventResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecretEventResponse", string(data)}, " ")
}
