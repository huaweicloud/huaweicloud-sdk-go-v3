package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Secret struct {

	// The secret identifier.
	SecretId string `json:"secret_id"`

	// The secret name.
	SecretName string `json:"secret_name"`
}

func (o Secret) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Secret struct{}"
	}

	return strings.Join([]string{"Secret", string(data)}, " ")
}
