package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetApplicationAssignmentConfigurationResponse Response Object
type GetApplicationAssignmentConfigurationResponse struct {

	// 应用程序是否需要分配
	AssignmentRequired *bool `json:"assignment_required,omitempty"`
	HttpStatusCode     int   `json:"-"`
}

func (o GetApplicationAssignmentConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetApplicationAssignmentConfigurationResponse struct{}"
	}

	return strings.Join([]string{"GetApplicationAssignmentConfigurationResponse", string(data)}, " ")
}
