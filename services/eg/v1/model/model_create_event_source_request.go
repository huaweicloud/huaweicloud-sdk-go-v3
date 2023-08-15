package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEventSourceRequest Request Object
type CreateEventSourceRequest struct {
	Body *CustomizeSourceCreateReq `json:"body,omitempty"`
}

func (o CreateEventSourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventSourceRequest struct{}"
	}

	return strings.Join([]string{"CreateEventSourceRequest", string(data)}, " ")
}
