package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckPublishAppRequest Request Object
type CheckPublishAppRequest struct {

	// 应用组ID。
	AppGroupId string `json:"app_group_id"`

	Body *CheckPublishAppReq `json:"body,omitempty"`
}

func (o CheckPublishAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPublishAppRequest struct{}"
	}

	return strings.Join([]string{"CheckPublishAppRequest", string(data)}, " ")
}
