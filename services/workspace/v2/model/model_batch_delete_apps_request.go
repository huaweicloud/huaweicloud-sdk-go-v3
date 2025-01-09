package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAppsRequest Request Object
type BatchDeleteAppsRequest struct {
	Body *BatchDeleteAppsReq `json:"body,omitempty"`
}

func (o BatchDeleteAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAppsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteAppsRequest", string(data)}, " ")
}
