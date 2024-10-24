package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgreementRequest Request Object
type ShowAgreementRequest struct {
}

func (o ShowAgreementRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgreementRequest struct{}"
	}

	return strings.Join([]string{"ShowAgreementRequest", string(data)}, " ")
}
