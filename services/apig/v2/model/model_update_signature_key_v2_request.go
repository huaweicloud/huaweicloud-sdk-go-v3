package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSignatureKeyV2Request Request Object
type UpdateSignatureKeyV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 签名密钥编号
	SignId string `json:"sign_id"`

	Body *BaseSignature `json:"body,omitempty"`
}

func (o UpdateSignatureKeyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSignatureKeyV2Request struct{}"
	}

	return strings.Join([]string{"UpdateSignatureKeyV2Request", string(data)}, " ")
}
