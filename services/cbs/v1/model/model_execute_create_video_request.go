package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExecuteCreateVideoRequest struct {
	Body *CreateVideoReq `json:"body,omitempty"`
}

func (o ExecuteCreateVideoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteCreateVideoRequest struct{}"
	}

	return strings.Join([]string{"ExecuteCreateVideoRequest", string(data)}, " ")
}
