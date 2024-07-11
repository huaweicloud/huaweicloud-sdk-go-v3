package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MoveAppGroupsRequest Request Object
type MoveAppGroupsRequest struct {

	// 项目Id
	ProjectId string `json:"project_id"`

	Body *MoveAppGroupsRequestBody `json:"body,omitempty"`
}

func (o MoveAppGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoveAppGroupsRequest struct{}"
	}

	return strings.Join([]string{"MoveAppGroupsRequest", string(data)}, " ")
}
