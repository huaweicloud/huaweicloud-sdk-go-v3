package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBuildJobResponseBodyResult 结果
type CreateBuildJobResponseBodyResult struct {

	// 构建任务ID
	JobId *string `json:"job_id,omitempty"`
}

func (o CreateBuildJobResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBuildJobResponseBodyResult struct{}"
	}

	return strings.Join([]string{"CreateBuildJobResponseBodyResult", string(data)}, " ")
}
