package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBuildJobResponseBodyResult 结果
type DeleteBuildJobResponseBodyResult struct {

	// 构建任务ID
	JobId *string `json:"job_id,omitempty"`

	// 构建任务所在项目的ID
	ProjectId *string `json:"project_id,omitempty"`
}

func (o DeleteBuildJobResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBuildJobResponseBodyResult struct{}"
	}

	return strings.Join([]string{"DeleteBuildJobResponseBodyResult", string(data)}, " ")
}
