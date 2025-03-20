package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAsymmetricSignatureReq Set asymmetric signature API request body.
type SetAsymmetricSignatureReq struct {
	AsymmetricSignature *AsymmetricSignature `json:"asymmetric_signature"`
}

func (o SetAsymmetricSignatureReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAsymmetricSignatureReq struct{}"
	}

	return strings.Join([]string{"SetAsymmetricSignatureReq", string(data)}, " ")
}
