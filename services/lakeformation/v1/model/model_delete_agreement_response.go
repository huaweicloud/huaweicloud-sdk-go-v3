package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAgreementResponse Response Object
type DeleteAgreementResponse struct {
	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAgreementResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAgreementResponse struct{}"
	}

	return strings.Join([]string{"DeleteAgreementResponse", string(data)}, " ")
}
