package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteSecretResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecretResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecretResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecretResponse", string(data)}, " ")
}
