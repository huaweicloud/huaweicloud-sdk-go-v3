package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateRomaAppRequest struct {
	// 应用ID

	AppId string `json:"app_id"`
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *UpdateAppReq `json:"body,omitempty"`
}

func (o UpdateRomaAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRomaAppRequest struct{}"
	}

	return strings.Join([]string{"UpdateRomaAppRequest", string(data)}, " ")
}
