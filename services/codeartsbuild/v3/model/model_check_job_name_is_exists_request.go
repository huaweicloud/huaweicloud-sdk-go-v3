package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckJobNameIsExistsRequest Request Object
type CheckJobNameIsExistsRequest struct {

	// CodeArts项目ID，32位数字、小写字母组合。
	ProjectId string `json:"project_id"`

	// 任务名称
	JobName string `json:"job_name"`
}

func (o CheckJobNameIsExistsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckJobNameIsExistsRequest struct{}"
	}

	return strings.Join([]string{"CheckJobNameIsExistsRequest", string(data)}, " ")
}
