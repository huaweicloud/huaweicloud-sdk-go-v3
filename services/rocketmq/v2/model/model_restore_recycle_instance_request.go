package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreRecycleInstanceRequest Request Object
type RestoreRecycleInstanceRequest struct {
	Body *BatchResumeInstanceReq `json:"body,omitempty"`
}

func (o RestoreRecycleInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreRecycleInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestoreRecycleInstanceRequest", string(data)}, " ")
}
