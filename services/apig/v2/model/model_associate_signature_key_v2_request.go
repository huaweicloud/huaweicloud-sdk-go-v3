package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AssociateSignatureKeyV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *SignBindingReq `json:"body,omitempty"`
}

func (o AssociateSignatureKeyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateSignatureKeyV2Request struct{}"
	}

	return strings.Join([]string{"AssociateSignatureKeyV2Request", string(data)}, " ")
}
