package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteV2AgreementRequest Request Object
type DeleteV2AgreementRequest struct {
}

func (o DeleteV2AgreementRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteV2AgreementRequest struct{}"
	}

	return strings.Join([]string{"DeleteV2AgreementRequest", string(data)}, " ")
}
