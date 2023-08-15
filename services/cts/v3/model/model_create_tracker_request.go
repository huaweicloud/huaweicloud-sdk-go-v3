package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTrackerRequest Request Object
type CreateTrackerRequest struct {
	Body *CreateTrackerRequestBody `json:"body,omitempty"`
}

func (o CreateTrackerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrackerRequest struct{}"
	}

	return strings.Join([]string{"CreateTrackerRequest", string(data)}, " ")
}
