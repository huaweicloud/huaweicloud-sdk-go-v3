package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSignatureKeyV2Request Request Object
type DeleteSignatureKeyV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 签名密钥编号
	SignId string `json:"sign_id"`
}

func (o DeleteSignatureKeyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSignatureKeyV2Request struct{}"
	}

	return strings.Join([]string{"DeleteSignatureKeyV2Request", string(data)}, " ")
}
