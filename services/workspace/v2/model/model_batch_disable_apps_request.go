package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDisableAppsRequest Request Object
type BatchDisableAppsRequest struct {
	Body *BatchOperateAppsReq `json:"body,omitempty"`
}

func (o BatchDisableAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisableAppsRequest struct{}"
	}

	return strings.Join([]string{"BatchDisableAppsRequest", string(data)}, " ")
}
