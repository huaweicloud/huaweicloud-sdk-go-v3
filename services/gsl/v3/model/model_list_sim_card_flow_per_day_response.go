package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSimCardFlowPerDayResponse Response Object
type ListSimCardFlowPerDayResponse struct {
	Body           *[]SimCardFlowPerDayRsp `json:"body,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListSimCardFlowPerDayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSimCardFlowPerDayResponse struct{}"
	}

	return strings.Join([]string{"ListSimCardFlowPerDayResponse", string(data)}, " ")
}
