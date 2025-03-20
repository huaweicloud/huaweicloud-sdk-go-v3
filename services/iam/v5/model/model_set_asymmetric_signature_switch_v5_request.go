package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAsymmetricSignatureSwitchV5Request Request Object
type SetAsymmetricSignatureSwitchV5Request struct {
	Body *SetAsymmetricSignatureReq `json:"body,omitempty"`
}

func (o SetAsymmetricSignatureSwitchV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAsymmetricSignatureSwitchV5Request struct{}"
	}

	return strings.Join([]string{"SetAsymmetricSignatureSwitchV5Request", string(data)}, " ")
}
