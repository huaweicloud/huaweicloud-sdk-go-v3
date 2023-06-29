package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Secret 密钥。
type Secret struct {
	Secret *SecretDetail `json:"secret"`
}

func (o Secret) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Secret struct{}"
	}

	return strings.Join([]string{"Secret", string(data)}, " ")
}
