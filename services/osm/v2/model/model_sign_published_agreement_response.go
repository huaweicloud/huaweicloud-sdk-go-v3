package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SignPublishedAgreementResponse Response Object
type SignPublishedAgreementResponse struct {

	// 签署记录id
	AgreementSignedRecordId *int32 `json:"agreement_signed_record_id,omitempty"`
	HttpStatusCode          int    `json:"-"`
}

func (o SignPublishedAgreementResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignPublishedAgreementResponse struct{}"
	}

	return strings.Join([]string{"SignPublishedAgreementResponse", string(data)}, " ")
}
