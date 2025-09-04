package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MoveGroupRequest Request Object
type MoveGroupRequest struct {

	// CodeArts项目ID，32位数字、小写字母组合。
	ProjectId string `json:"project_id"`

	Body *MoveJobGroupRequestBody `json:"body,omitempty"`
}

func (o MoveGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoveGroupRequest struct{}"
	}

	return strings.Join([]string{"MoveGroupRequest", string(data)}, " ")
}
