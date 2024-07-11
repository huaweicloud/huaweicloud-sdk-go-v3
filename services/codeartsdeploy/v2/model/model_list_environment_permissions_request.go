package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnvironmentPermissionsRequest Request Object
type ListEnvironmentPermissionsRequest struct {

	// 应用id
	ApplicationId string `json:"application_id"`

	// 环境id
	EnvironmentId string `json:"environment_id"`
}

func (o ListEnvironmentPermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvironmentPermissionsRequest struct{}"
	}

	return strings.Join([]string{"ListEnvironmentPermissionsRequest", string(data)}, " ")
}
