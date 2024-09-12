package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSimCardFlowPerDayRequest Request Object
type ListSimCardFlowPerDayRequest struct {
	Body *SimCardFlowPerDayReq `json:"body,omitempty"`
}

func (o ListSimCardFlowPerDayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSimCardFlowPerDayRequest struct{}"
	}

	return strings.Join([]string{"ListSimCardFlowPerDayRequest", string(data)}, " ")
}
