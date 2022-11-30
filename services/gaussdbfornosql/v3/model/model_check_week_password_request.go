package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckWeekPasswordRequest struct {
	Body *CheckWeekPasswordRequestBody `json:"body,omitempty"`
}

func (o CheckWeekPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckWeekPasswordRequest struct{}"
	}

	return strings.Join([]string{"CheckWeekPasswordRequest", string(data)}, " ")
}
