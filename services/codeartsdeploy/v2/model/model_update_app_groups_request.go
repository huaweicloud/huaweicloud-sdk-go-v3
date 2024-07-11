package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppGroupsRequest Request Object
type UpdateAppGroupsRequest struct {

	// 项目Id
	ProjectId string `json:"project_id"`

	// 分组Id
	GroupId string `json:"group_id"`

	Body *UpdateAppGroupsRequestBody `json:"body,omitempty"`
}

func (o UpdateAppGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppGroupsRequest struct{}"
	}

	return strings.Join([]string{"UpdateAppGroupsRequest", string(data)}, " ")
}
