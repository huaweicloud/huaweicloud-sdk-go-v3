package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetAsymmetricSignatureSwitchV5Request Request Object
type GetAsymmetricSignatureSwitchV5Request struct {
}

func (o GetAsymmetricSignatureSwitchV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAsymmetricSignatureSwitchV5Request struct{}"
	}

	return strings.Join([]string{"GetAsymmetricSignatureSwitchV5Request", string(data)}, " ")
}
