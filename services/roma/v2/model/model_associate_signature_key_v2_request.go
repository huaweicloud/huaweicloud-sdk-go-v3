package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateSignatureKeyV2Request Request Object
type AssociateSignatureKeyV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *SignApiBinding `json:"body,omitempty"`
}

func (o AssociateSignatureKeyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateSignatureKeyV2Request struct{}"
	}

	return strings.Join([]string{"AssociateSignatureKeyV2Request", string(data)}, " ")
}
