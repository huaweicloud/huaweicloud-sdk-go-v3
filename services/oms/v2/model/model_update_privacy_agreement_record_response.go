package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePrivacyAgreementRecordResponse Response Object
type UpdatePrivacyAgreementRecordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePrivacyAgreementRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivacyAgreementRecordResponse struct{}"
	}

	return strings.Join([]string{"UpdatePrivacyAgreementRecordResponse", string(data)}, " ")
}
