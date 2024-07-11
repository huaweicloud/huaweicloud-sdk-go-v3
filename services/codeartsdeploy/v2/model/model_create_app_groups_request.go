package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppGroupsRequest Request Object
type CreateAppGroupsRequest struct {

	// 项目Id
	ProjectId string `json:"project_id"`

	Body *CreateAppGroupsRequestBody `json:"body,omitempty"`
}

func (o CreateAppGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppGroupsRequest struct{}"
	}

	return strings.Join([]string{"CreateAppGroupsRequest", string(data)}, " ")
}
