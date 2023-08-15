package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckPutEventsRequest Request Object
type CheckPutEventsRequest struct {
	Body *CheckPutEventsReq `json:"body,omitempty"`
}

func (o CheckPutEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPutEventsRequest struct{}"
	}

	return strings.Join([]string{"CheckPutEventsRequest", string(data)}, " ")
}
