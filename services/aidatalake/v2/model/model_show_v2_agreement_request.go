package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowV2AgreementRequest Request Object
type ShowV2AgreementRequest struct {
}

func (o ShowV2AgreementRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowV2AgreementRequest struct{}"
	}

	return strings.Join([]string{"ShowV2AgreementRequest", string(data)}, " ")
}
