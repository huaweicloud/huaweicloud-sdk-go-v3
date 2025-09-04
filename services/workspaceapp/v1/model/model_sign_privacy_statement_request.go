package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SignPrivacyStatementRequest Request Object
type SignPrivacyStatementRequest struct {
	Body *SignPrivacyStatementReq `json:"body,omitempty"`
}

func (o SignPrivacyStatementRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignPrivacyStatementRequest struct{}"
	}

	return strings.Join([]string{"SignPrivacyStatementRequest", string(data)}, " ")
}
