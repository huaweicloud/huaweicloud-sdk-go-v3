package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCaptureResultResponse Response Object
type ListCaptureResultResponse struct {
	Data           *CaptureResultUrlVo `json:"data,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListCaptureResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCaptureResultResponse struct{}"
	}

	return strings.Join([]string{"ListCaptureResultResponse", string(data)}, " ")
}
