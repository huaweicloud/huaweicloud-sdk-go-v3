package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type UpdateResOnlineInstanceRequestBody struct {

	// 类别。
	Category string `json:"category" xml:"category"`

	// 描述。
	Description *string `json:"description,omitempty" xml:"description"`

	JobConfig *JobConfig `json:"job_config" xml:"job_config"`

	// 作业名称。
	JobName string `json:"job_name" xml:"job_name"`

	// 作业类型。
	JobType string `json:"job_type" xml:"job_type"`

	// 调度参数。
	Schedule string `json:"schedule" xml:"schedule"`
}

func (o UpdateResOnlineInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResOnlineInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateResOnlineInstanceRequestBody", string(data)}, " ")
}
