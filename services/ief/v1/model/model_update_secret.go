package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecret 密钥。
type UpdateSecret struct {
	Secret *UpdateSecretDetail `json:"secret"`
}

func (o UpdateSecret) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecret struct{}"
	}

	return strings.Join([]string{"UpdateSecret", string(data)}, " ")
}
