package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTokenVerificationRequest struct {
}

func (o ShowTokenVerificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTokenVerificationRequest struct{}"
	}

	return strings.Join([]string{"ShowTokenVerificationRequest", string(data)}, " ")
}
