package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTheJobResponseBodyResult 结果
type DeleteTheJobResponseBodyResult struct {

	// 构建任务ID
	JobId *string `json:"job_id,omitempty"`

	// 构建任务所在项目的ID
	ProjectId *string `json:"project_id,omitempty"`
}

func (o DeleteTheJobResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTheJobResponseBodyResult struct{}"
	}

	return strings.Join([]string{"DeleteTheJobResponseBodyResult", string(data)}, " ")
}
