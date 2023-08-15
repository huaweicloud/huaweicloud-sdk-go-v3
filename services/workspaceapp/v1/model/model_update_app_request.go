package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppRequest Request Object
type UpdateAppRequest struct {

	// 应用组ID
	AppGroupId string `json:"app_group_id"`

	// 应用ID
	AppId string `json:"app_id"`

	Body *UpdateAppReq `json:"body,omitempty"`
}

func (o UpdateAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppRequest struct{}"
	}

	return strings.Join([]string{"UpdateAppRequest", string(data)}, " ")
}
