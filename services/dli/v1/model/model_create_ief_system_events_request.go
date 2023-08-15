package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIefSystemEventsRequest Request Object
type CreateIefSystemEventsRequest struct {
	Body *IefSystemEventsReq `json:"body,omitempty"`
}

func (o CreateIefSystemEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIefSystemEventsRequest struct{}"
	}

	return strings.Join([]string{"CreateIefSystemEventsRequest", string(data)}, " ")
}
