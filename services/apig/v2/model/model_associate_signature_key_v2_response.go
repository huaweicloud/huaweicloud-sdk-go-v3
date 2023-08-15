package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateSignatureKeyV2Response Response Object
type AssociateSignatureKeyV2Response struct {

	// API与签名密钥的绑定关系列表
	Bindings       *[]SignApiBindingInfo `json:"bindings,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o AssociateSignatureKeyV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateSignatureKeyV2Response struct{}"
	}

	return strings.Join([]string{"AssociateSignatureKeyV2Response", string(data)}, " ")
}
