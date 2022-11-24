package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateSignatureKeyV2Request struct {

	// 实例ID
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
