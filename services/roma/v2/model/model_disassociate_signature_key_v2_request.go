package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DisassociateSignatureKeyV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// API与签名密钥的绑定关系编号

	SignBindingsId string `json:"sign_bindings_id"`
}

func (o DisassociateSignatureKeyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateSignatureKeyV2Request struct{}"
	}

	return strings.Join([]string{"DisassociateSignatureKeyV2Request", string(data)}, " ")
}
