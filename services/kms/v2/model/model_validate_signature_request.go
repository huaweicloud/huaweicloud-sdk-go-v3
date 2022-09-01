package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ValidateSignatureRequest struct {
	Body *VerifyRequestBody `json:"body,omitempty" xml:"body"`
}

func (o ValidateSignatureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateSignatureRequest struct{}"
	}

	return strings.Join([]string{"ValidateSignatureRequest", string(data)}, " ")
}
