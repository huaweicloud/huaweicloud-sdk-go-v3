package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAimMsgSignatureRequest Request Object
type AddAimMsgSignatureRequest struct {
	Body *SignatureRequest `json:"body,omitempty"`
}

func (o AddAimMsgSignatureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAimMsgSignatureRequest struct{}"
	}

	return strings.Join([]string{"AddAimMsgSignatureRequest", string(data)}, " ")
}
