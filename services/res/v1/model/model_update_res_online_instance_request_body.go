package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResOnlineInstanceRequestBody This is a auto create Body Object
type UpdateResOnlineInstanceRequestBody struct {

	// 类别。
	Category string `json:"category"`

	// 描述。
	Description *string `json:"description,omitempty"`

	JobConfig *JobConfig `json:"job_config"`

	// 作业名称。
	JobName string `json:"job_name"`

	// 作业类型。
	JobType string `json:"job_type"`

	// 调度参数。
	Schedule string `json:"schedule"`
}

func (o UpdateResOnlineInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResOnlineInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateResOnlineInstanceRequestBody", string(data)}, " ")
}
