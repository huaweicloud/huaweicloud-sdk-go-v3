package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type PauseResumeDataSynchronizationResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PauseResumeDataSynchronizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseResumeDataSynchronizationResponse struct{}"
	}

	return strings.Join([]string{"PauseResumeDataSynchronizationResponse", string(data)}, " ")
}
