package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGroupRequest Request Object
type CreateGroupRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`

	Body *CreateGroupRequestBody `json:"body,omitempty"`
}

func (o CreateGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateGroupRequest", string(data)}, " ")
}
