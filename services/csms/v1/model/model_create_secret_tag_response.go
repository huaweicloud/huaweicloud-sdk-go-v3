package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecretTagResponse Response Object
type CreateSecretTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateSecretTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecretTagResponse struct{}"
	}

	return strings.Join([]string{"CreateSecretTagResponse", string(data)}, " ")
}
