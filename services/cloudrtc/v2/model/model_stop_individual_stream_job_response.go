package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopIndividualStreamJobResponse Response Object
type StopIndividualStreamJobResponse struct {
	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopIndividualStreamJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopIndividualStreamJobResponse struct{}"
	}

	return strings.Join([]string{"StopIndividualStreamJobResponse", string(data)}, " ")
}
