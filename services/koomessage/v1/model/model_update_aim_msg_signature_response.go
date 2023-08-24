package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAimMsgSignatureResponse Response Object
type UpdateAimMsgSignatureResponse struct {

	// 签名ID。
	SignatureId *string `json:"signature_id,omitempty"`

	// 签名名称。
	SignatureName  *string `json:"signature_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAimMsgSignatureResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAimMsgSignatureResponse struct{}"
	}

	return strings.Join([]string{"UpdateAimMsgSignatureResponse", string(data)}, " ")
}
