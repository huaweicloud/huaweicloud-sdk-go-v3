package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSignatureRequest Request Object
type CreateSignatureRequest struct {
	Body *SmsSignatureReq `json:"body,omitempty"`
}

func (o CreateSignatureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSignatureRequest struct{}"
	}

	return strings.Join([]string{"CreateSignatureRequest", string(data)}, " ")
}
