package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDomainVerificationRequest Request Object
type CheckDomainVerificationRequest struct {
	Body *CheckDomainVerificationInfo `json:"body,omitempty"`
}

func (o CheckDomainVerificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDomainVerificationRequest struct{}"
	}

	return strings.Join([]string{"CheckDomainVerificationRequest", string(data)}, " ")
}
