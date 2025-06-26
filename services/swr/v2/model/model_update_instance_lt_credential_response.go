package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceLtCredentialResponse Response Object
type UpdateInstanceLtCredentialResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceLtCredentialResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceLtCredentialResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceLtCredentialResponse", string(data)}, " ")
}
