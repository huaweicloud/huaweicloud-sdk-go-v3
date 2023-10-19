package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAgreementRequest Request Object
type DeleteAgreementRequest struct {
}

func (o DeleteAgreementRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAgreementRequest struct{}"
	}

	return strings.Join([]string{"DeleteAgreementRequest", string(data)}, " ")
}
