package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 最近作业详情
type JobRsp struct {

	// 项目名称
	EihealthProjectName *string `json:"eihealth_project_name,omitempty"`

	// 项目id
	EihealthProjectId *string `json:"eihealth_project_id,omitempty"`

	// 作业名称
	Name *string `json:"name,omitempty"`

	// 作业id
	Id *string `json:"id,omitempty"`

	// 作业状态
	Status *string `json:"status,omitempty"`

	// 作业结束时间
	FinishTime *string `json:"finish_time,omitempty"`

	// 作业失败提示
	FailedMessage *string `json:"failed_message,omitempty"`

	// 作业失败原因
	FailedReason *string `json:"failed_reason,omitempty"`
}

func (o JobRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobRsp struct{}"
	}

	return strings.Join([]string{"JobRsp", string(data)}, " ")
}
