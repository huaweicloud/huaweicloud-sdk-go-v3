package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkspacesForUserRequest Request Object
type ListWorkspacesForUserRequest struct {

	// DataArtsStudio实例id
	InstanceId string `json:"instance_id"`

	// 用户id
	UserId string `json:"user_id"`
}

func (o ListWorkspacesForUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkspacesForUserRequest struct{}"
	}

	return strings.Join([]string{"ListWorkspacesForUserRequest", string(data)}, " ")
}
