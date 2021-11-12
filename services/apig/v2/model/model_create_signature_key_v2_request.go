package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateSignatureKeyV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *SignatureReq `json:"body,omitempty"`
}

func (o CreateSignatureKeyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSignatureKeyV2Request struct{}"
	}

	return strings.Join([]string{"CreateSignatureKeyV2Request", string(data)}, " ")
}
