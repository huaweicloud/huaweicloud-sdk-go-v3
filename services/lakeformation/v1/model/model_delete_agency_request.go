package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAgencyRequest Request Object
type DeleteAgencyRequest struct {
	Body *AgencyRequestBody `json:"body,omitempty"`
}

func (o DeleteAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAgencyRequest struct{}"
	}

	return strings.Join([]string{"DeleteAgencyRequest", string(data)}, " ")
}
