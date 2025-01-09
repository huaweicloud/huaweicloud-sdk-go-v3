package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchEnableAppRequest Request Object
type BatchEnableAppRequest struct {

	// 应用组ID。
	AppGroupId string `json:"app_group_id"`

	Body *BatchUpdateAppReq `json:"body,omitempty"`
}

func (o BatchEnableAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchEnableAppRequest struct{}"
	}

	return strings.Join([]string{"BatchEnableAppRequest", string(data)}, " ")
}
