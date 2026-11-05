package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VerifyAgencyRequest Request Object
type VerifyAgencyRequest struct {
	Body *AgencyRequest `json:"body,omitempty"`
}

func (o VerifyAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerifyAgencyRequest struct{}"
	}

	return strings.Join([]string{"VerifyAgencyRequest", string(data)}, " ")
}
