package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPrivacyAgreementsRequest Request Object
type ShowPrivacyAgreementsRequest struct {
}

func (o ShowPrivacyAgreementsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivacyAgreementsRequest struct{}"
	}

	return strings.Join([]string{"ShowPrivacyAgreementsRequest", string(data)}, " ")
}
