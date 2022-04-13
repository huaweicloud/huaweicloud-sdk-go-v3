package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteSignatureKeyV2Request struct {
	// 实例ID

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
