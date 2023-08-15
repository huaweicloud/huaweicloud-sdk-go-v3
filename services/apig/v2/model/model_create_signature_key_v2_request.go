package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSignatureKeyV2Request Request Object
type CreateSignatureKeyV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *BaseSignature `json:"body,omitempty"`
}

func (o CreateSignatureKeyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSignatureKeyV2Request struct{}"
	}

	return strings.Join([]string{"CreateSignatureKeyV2Request", string(data)}, " ")
}
