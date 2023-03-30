package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateApplyResponse struct {

	// the message of join-request
	Message *string `json:"message,omitempty"`

	// the status of join-request(success/error/conflict)
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateApplyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplyResponse struct{}"
	}

	return strings.Join([]string{"CreateApplyResponse", string(data)}, " ")
}
