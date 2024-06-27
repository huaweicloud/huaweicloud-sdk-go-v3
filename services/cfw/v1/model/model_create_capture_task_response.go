package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCaptureTaskResponse Response Object
type CreateCaptureTaskResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateCaptureTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCaptureTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateCaptureTaskResponse", string(data)}, " ")
}
