package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTrackerRequest Request Object
type UpdateTrackerRequest struct {
	Body *UpdateTrackerRequestBody `json:"body,omitempty"`
}

func (o UpdateTrackerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTrackerRequest struct{}"
	}

	return strings.Join([]string{"UpdateTrackerRequest", string(data)}, " ")
}
