package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecyclingJob struct {

	// 任务ID
	JobId *string `json:"job_id,omitempty"`

	// 任务名称
	JobName *string `json:"job_name,omitempty"`

	// CodeArts项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 创建人
	CreatorName *string `json:"creator_name,omitempty"`

	// 创建人ID
	CreatorId *string `json:"creator_id,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 删除时间
	DeleteTime *string `json:"delete_time,omitempty"`

	// 创建时间戳
	CreateTimeStamp *string `json:"create_time_stamp,omitempty"`

	// 删除时间戳
	DeleteTimeStamp *string `json:"delete_time_stamp,omitempty"`
}

func (o RecyclingJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecyclingJob struct{}"
	}

	return strings.Join([]string{"RecyclingJob", string(data)}, " ")
}
