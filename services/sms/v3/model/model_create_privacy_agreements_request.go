package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrivacyAgreementsRequest Request Object
type CreatePrivacyAgreementsRequest struct {
}

func (o CreatePrivacyAgreementsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivacyAgreementsRequest struct{}"
	}

	return strings.Join([]string{"CreatePrivacyAgreementsRequest", string(data)}, " ")
}
