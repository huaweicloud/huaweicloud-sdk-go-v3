package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSignedLatestPublishedAgreementResponse Response Object
type ShowSignedLatestPublishedAgreementResponse struct {

	// 是否签署
	IsSignedLatest *bool `json:"is_signed_latest,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowSignedLatestPublishedAgreementResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSignedLatestPublishedAgreementResponse struct{}"
	}

	return strings.Join([]string{"ShowSignedLatestPublishedAgreementResponse", string(data)}, " ")
}
