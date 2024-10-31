package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCaptureTaskResponse Response Object
type ListCaptureTaskResponse struct {
	Data           *HttpQueryCaptureTaskResponseData `json:"data,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListCaptureTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCaptureTaskResponse struct{}"
	}

	return strings.Join([]string{"ListCaptureTaskResponse", string(data)}, " ")
}
