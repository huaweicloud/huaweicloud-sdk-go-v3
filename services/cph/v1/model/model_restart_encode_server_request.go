package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RestartEncodeServerRequest struct {
	Body *RestartEncodeServerRequestBody `json:"body,omitempty"`
}

func (o RestartEncodeServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartEncodeServerRequest struct{}"
	}

	return strings.Join([]string{"RestartEncodeServerRequest", string(data)}, " ")
}
