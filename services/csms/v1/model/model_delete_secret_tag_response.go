package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecretTagResponse Response Object
type DeleteSecretTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecretTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecretTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecretTagResponse", string(data)}, " ")
}
