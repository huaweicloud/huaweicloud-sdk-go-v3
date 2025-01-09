package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchEnableAppsRequest Request Object
type BatchEnableAppsRequest struct {
	Body *BatchOperateAppsReq `json:"body,omitempty"`
}

func (o BatchEnableAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchEnableAppsRequest struct{}"
	}

	return strings.Join([]string{"BatchEnableAppsRequest", string(data)}, " ")
}
