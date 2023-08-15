package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSignatureResponse Response Object
type CreateSignatureResponse struct {

	// 签名主键ID
	Id *string `json:"id,omitempty"`

	// 签名ID
	SignatureId *string `json:"signature_id,omitempty"`

	// 签名名称
	SignatureName  *string `json:"signature_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSignatureResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSignatureResponse struct{}"
	}

	return strings.Join([]string{"CreateSignatureResponse", string(data)}, " ")
}
