package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimMsgSignatureDetailRequest Request Object
type ListAimMsgSignatureDetailRequest struct {

	// 签名ID。
	SignatureId string `json:"signature_id"`
}

func (o ListAimMsgSignatureDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimMsgSignatureDetailRequest struct{}"
	}

	return strings.Join([]string{"ListAimMsgSignatureDetailRequest", string(data)}, " ")
}
