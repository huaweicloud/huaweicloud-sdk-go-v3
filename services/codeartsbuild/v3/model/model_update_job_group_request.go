package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateJobGroupRequest Request Object
type UpdateJobGroupRequest struct {

	// CodeArts项目ID，32位数字、小写字母组合。
	ProjectId string `json:"project_id"`

	Body *JobGroupRequestBody `json:"body,omitempty"`
}

func (o UpdateJobGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateJobGroupRequest", string(data)}, " ")
}
