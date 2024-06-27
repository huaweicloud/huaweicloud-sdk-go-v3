package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelCaptureTaskResponse Response Object
type CancelCaptureTaskResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CancelCaptureTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelCaptureTaskResponse struct{}"
	}

	return strings.Join([]string{"CancelCaptureTaskResponse", string(data)}, " ")
}
