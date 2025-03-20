package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAsymmetricSignatureSwitchV5Response Response Object
type SetAsymmetricSignatureSwitchV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o SetAsymmetricSignatureSwitchV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAsymmetricSignatureSwitchV5Response struct{}"
	}

	return strings.Join([]string{"SetAsymmetricSignatureSwitchV5Response", string(data)}, " ")
}
