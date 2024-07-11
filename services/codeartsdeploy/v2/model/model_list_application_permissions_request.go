package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationPermissionsRequest Request Object
type ListApplicationPermissionsRequest struct {

	// 应用id
	AppId *string `json:"app_id,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`
}

func (o ListApplicationPermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationPermissionsRequest struct{}"
	}

	return strings.Join([]string{"ListApplicationPermissionsRequest", string(data)}, " ")
}
