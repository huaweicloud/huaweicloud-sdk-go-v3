package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MemberCheckJob 参数解释：后端服务器检测任务信息。
type MemberCheckJob struct {

	// 参数解释：任务ID，根据该任务ID可以查询到任务执行结果。
	JobId *string `json:"job_id,omitempty"`
}

func (o MemberCheckJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberCheckJob struct{}"
	}

	return strings.Join([]string{"MemberCheckJob", string(data)}, " ")
}
