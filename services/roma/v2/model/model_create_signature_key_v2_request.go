package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSignatureKeyV2Request Request Object
type CreateSignatureKeyV2Request struct {

	// 实例ID
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
