package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSignatureRequest Request Object
type UpdateSignatureRequest struct {

	// 签名ID
	Id string `json:"id"`

	Body *SmsSignatureReq `json:"body,omitempty"`
}

func (o UpdateSignatureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSignatureRequest struct{}"
	}

	return strings.Join([]string{"UpdateSignatureRequest", string(data)}, " ")
}
