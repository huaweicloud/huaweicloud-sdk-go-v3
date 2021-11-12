package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DisassociateSignatureKeyV2Request struct {
	// 实例编号

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
