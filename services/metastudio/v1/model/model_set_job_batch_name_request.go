package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetJobBatchNameRequest Request Object
type SetJobBatchNameRequest struct {
	Body *SetJobBatchNameReq `json:"body,omitempty"`
}

func (o SetJobBatchNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetJobBatchNameRequest struct{}"
	}

	return strings.Join([]string{"SetJobBatchNameRequest", string(data)}, " ")
}
