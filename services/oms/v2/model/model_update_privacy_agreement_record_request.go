package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePrivacyAgreementRecordRequest Request Object
type UpdatePrivacyAgreementRecordRequest struct {
}

func (o UpdatePrivacyAgreementRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivacyAgreementRecordRequest struct{}"
	}

	return strings.Join([]string{"UpdatePrivacyAgreementRecordRequest", string(data)}, " ")
}
