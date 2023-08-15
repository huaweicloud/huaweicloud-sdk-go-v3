package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnpublishAppRequest Request Object
type UnpublishAppRequest struct {

	// 应用组ID
	AppGroupId string `json:"app_group_id"`

	Body *UnpublishAppReq `json:"body,omitempty"`
}

func (o UnpublishAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnpublishAppRequest struct{}"
	}

	return strings.Join([]string{"UnpublishAppRequest", string(data)}, " ")
}
