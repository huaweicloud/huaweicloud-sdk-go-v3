package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceLtCredentialResponse Response Object
type DeleteInstanceLtCredentialResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInstanceLtCredentialResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceLtCredentialResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceLtCredentialResponse", string(data)}, " ")
}
