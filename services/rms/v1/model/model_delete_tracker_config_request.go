package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteTrackerConfigRequest struct {
}

func (o DeleteTrackerConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrackerConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteTrackerConfigRequest", string(data)}, " ")
}
