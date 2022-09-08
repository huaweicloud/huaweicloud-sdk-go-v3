package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StopCloudPhoneRequest struct {
	Body *StopCloudPhoneRequestBody `json:"body,omitempty"`
}

func (o StopCloudPhoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopCloudPhoneRequest struct{}"
	}

	return strings.Join([]string{"StopCloudPhoneRequest", string(data)}, " ")
}
