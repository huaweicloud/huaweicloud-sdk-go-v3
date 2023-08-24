package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAimMsgSignatureRequest Request Object
type UpdateAimMsgSignatureRequest struct {

	// 签名ID。
	SignatureId string `json:"signature_id"`

	Body *SignatureRequest `json:"body,omitempty"`
}

func (o UpdateAimMsgSignatureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAimMsgSignatureRequest struct{}"
	}

	return strings.Join([]string{"UpdateAimMsgSignatureRequest", string(data)}, " ")
}
