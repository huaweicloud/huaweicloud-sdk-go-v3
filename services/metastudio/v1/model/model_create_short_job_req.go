package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateShortJobReq 创建短任务请求。
type CreateShortJobReq struct {
	JobType *ShortJobType `json:"job_type"`

	// 语言。 * CN: 中文 * EN: 英文
	Language *string `json:"language,omitempty"`
}

func (o CreateShortJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShortJobReq struct{}"
	}

	return strings.Join([]string{"CreateShortJobReq", string(data)}, " ")
}
