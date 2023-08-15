package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishAppRequest Request Object
type PublishAppRequest struct {

	// 应用组ID
	AppGroupId string `json:"app_group_id"`

	Body *PublishAppReq `json:"body,omitempty"`
}

func (o PublishAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishAppRequest struct{}"
	}

	return strings.Join([]string{"PublishAppRequest", string(data)}, " ")
}
