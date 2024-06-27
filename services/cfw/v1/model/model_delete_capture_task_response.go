package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCaptureTaskResponse Response Object
type DeleteCaptureTaskResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeleteCaptureTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCaptureTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteCaptureTaskResponse", string(data)}, " ")
}
