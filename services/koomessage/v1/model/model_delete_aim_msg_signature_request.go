package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAimMsgSignatureRequest Request Object
type DeleteAimMsgSignatureRequest struct {

	// 签名ID。
	SignatureId string `json:"signature_id"`
}

func (o DeleteAimMsgSignatureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAimMsgSignatureRequest struct{}"
	}

	return strings.Join([]string{"DeleteAimMsgSignatureRequest", string(data)}, " ")
}
