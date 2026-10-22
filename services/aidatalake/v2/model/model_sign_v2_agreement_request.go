package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SignV2AgreementRequest Request Object
type SignV2AgreementRequest struct {
	Body *SignV2AgreementRequestBody `json:"body,omitempty"`
}

func (o SignV2AgreementRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignV2AgreementRequest struct{}"
	}

	return strings.Join([]string{"SignV2AgreementRequest", string(data)}, " ")
}
