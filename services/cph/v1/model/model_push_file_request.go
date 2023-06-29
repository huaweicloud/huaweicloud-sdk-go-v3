package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PushFileRequest Request Object
type PushFileRequest struct {
	Body *PushFileRequestBody `json:"body,omitempty"`
}

func (o PushFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushFileRequest struct{}"
	}

	return strings.Join([]string{"PushFileRequest", string(data)}, " ")
}
