package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopSmartLiveResponse Response Object
type StopSmartLiveResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopSmartLiveResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopSmartLiveResponse struct{}"
	}

	return strings.Join([]string{"StopSmartLiveResponse", string(data)}, " ")
}
