package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckTokenVerificationRequest Request Object
type CheckTokenVerificationRequest struct {
}

func (o CheckTokenVerificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckTokenVerificationRequest struct{}"
	}

	return strings.Join([]string{"CheckTokenVerificationRequest", string(data)}, " ")
}
