package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunSyncCommandRequest Request Object
type RunSyncCommandRequest struct {
	Body *RunSyncCommandRequestBody `json:"body,omitempty"`
}

func (o RunSyncCommandRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSyncCommandRequest struct{}"
	}

	return strings.Join([]string{"RunSyncCommandRequest", string(data)}, " ")
}
