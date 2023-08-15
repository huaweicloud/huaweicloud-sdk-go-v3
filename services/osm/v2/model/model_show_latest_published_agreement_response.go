package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLatestPublishedAgreementResponse Response Object
type ShowLatestPublishedAgreementResponse struct {
	CaseAgreement  *AgreementVo `json:"case_agreement,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowLatestPublishedAgreementResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLatestPublishedAgreementResponse struct{}"
	}

	return strings.Join([]string{"ShowLatestPublishedAgreementResponse", string(data)}, " ")
}
