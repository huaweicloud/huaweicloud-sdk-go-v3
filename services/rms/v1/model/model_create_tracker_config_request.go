package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateTrackerConfigRequest struct {
	Body *TrackerConfigBody `json:"body,omitempty" xml:"body"`
}

func (o CreateTrackerConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrackerConfigRequest struct{}"
	}

	return strings.Join([]string{"CreateTrackerConfigRequest", string(data)}, " ")
}
