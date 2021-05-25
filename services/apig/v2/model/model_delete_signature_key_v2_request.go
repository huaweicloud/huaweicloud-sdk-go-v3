package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteSignatureKeyV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 签名密钥编号

	SignId string `json:"sign_id"`
}

func (o DeleteSignatureKeyV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSignatureKeyV2Request struct{}"
	}

	return strings.Join([]string{"DeleteSignatureKeyV2Request", string(data)}, " ")
}
