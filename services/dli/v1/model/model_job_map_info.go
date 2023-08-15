package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobMapInfo 导入job作业结果的模型。
type JobMapInfo struct {

	// 导入文件中的作业ID。
	OldJobId *int64 `json:"old_job_id,omitempty"`

	// 导入完成后作业ID，若is_cover=false，服务中已有同名的作业，则为-1。
	NewJobId *int64 `json:"new_job_id,omitempty"`

	// 导入作业结果信息。
	Remark *string `json:"remark,omitempty"`
}

func (o JobMapInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobMapInfo struct{}"
	}

	return strings.Join([]string{"JobMapInfo", string(data)}, " ")
}
