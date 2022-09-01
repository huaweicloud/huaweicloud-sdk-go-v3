package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteDeploymentGroupResponse struct {

	// 主机组ID
	GroupId        *string `json:"group_id,omitempty" xml:"group_id"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDeploymentGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeploymentGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteDeploymentGroupResponse", string(data)}, " ")
}
