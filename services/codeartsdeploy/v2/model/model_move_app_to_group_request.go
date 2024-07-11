package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MoveAppToGroupRequest Request Object
type MoveAppToGroupRequest struct {

	// 项目Id
	ProjectId string `json:"project_id"`

	Body *MoveAppToGroupRequestBody `json:"body,omitempty"`
}

func (o MoveAppToGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoveAppToGroupRequest struct{}"
	}

	return strings.Join([]string{"MoveAppToGroupRequest", string(data)}, " ")
}
