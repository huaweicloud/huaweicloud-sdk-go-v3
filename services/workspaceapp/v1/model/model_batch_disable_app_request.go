package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDisableAppRequest Request Object
type BatchDisableAppRequest struct {

	// 应用组ID。
	AppGroupId string `json:"app_group_id"`

	Body *BatchUpdateAppReq `json:"body,omitempty"`
}

func (o BatchDisableAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisableAppRequest struct{}"
	}

	return strings.Join([]string{"BatchDisableAppRequest", string(data)}, " ")
}
