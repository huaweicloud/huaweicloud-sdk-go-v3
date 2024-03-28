package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIefSystemEventsRequest Request Object
type CreateIefSystemEventsRequest struct {
	Body *CreateIefSystemEventsRequestBody `json:"body,omitempty"`
}

func (o CreateIefSystemEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIefSystemEventsRequest struct{}"
	}

	return strings.Join([]string{"CreateIefSystemEventsRequest", string(data)}, " ")
}
