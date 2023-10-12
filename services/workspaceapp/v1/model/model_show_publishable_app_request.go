package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublishableAppRequest Request Object
type ShowPublishableAppRequest struct {

	// 应用组ID
	AppGroupId string `json:"app_group_id"`
}

func (o ShowPublishableAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublishableAppRequest struct{}"
	}

	return strings.Join([]string{"ShowPublishableAppRequest", string(data)}, " ")
}
