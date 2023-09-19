package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrivacyAgreementsResponse Response Object
type CreatePrivacyAgreementsResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePrivacyAgreementsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivacyAgreementsResponse struct{}"
	}

	return strings.Join([]string{"CreatePrivacyAgreementsResponse", string(data)}, " ")
}
