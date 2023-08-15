package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopMixJobResponse Response Object
type StopMixJobResponse struct {
	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopMixJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopMixJobResponse struct{}"
	}

	return strings.Join([]string{"StopMixJobResponse", string(data)}, " ")
}
