package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddUserToAppRequest struct {
	// 应用ID

	AppId string `json:"app_id"`
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *AddUserToApp `json:"body,omitempty"`
}

func (o AddUserToAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddUserToAppRequest struct{}"
	}

	return strings.Join([]string{"AddUserToAppRequest", string(data)}, " ")
}
