package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateSignatureResponse struct {

	// 签名主键ID
	Id *string `json:"id,omitempty"`

	// 签名ID
	SignatureId *string `json:"signature_id,omitempty"`

	// 签名名称
	SignatureName  *string `json:"signature_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSignatureResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSignatureResponse struct{}"
	}

	return strings.Join([]string{"UpdateSignatureResponse", string(data)}, " ")
}
