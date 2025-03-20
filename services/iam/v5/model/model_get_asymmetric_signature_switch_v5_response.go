package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetAsymmetricSignatureSwitchV5Response Response Object
type GetAsymmetricSignatureSwitchV5Response struct {
	AsymmetricSignature *AsymmetricSignatureWithDomainId `json:"asymmetric_signature,omitempty"`
	HttpStatusCode      int                              `json:"-"`
}

func (o GetAsymmetricSignatureSwitchV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAsymmetricSignatureSwitchV5Response struct{}"
	}

	return strings.Join([]string{"GetAsymmetricSignatureSwitchV5Response", string(data)}, " ")
}
