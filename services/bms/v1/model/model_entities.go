package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Entities entities字段数据结构说明
type Entities struct {

	// 裸金属服务器相关操作显示server_id
	ServerId *string `json:"server_id,omitempty"`

	// 网卡相关操作显示nic_id
	NicId *string `json:"nic_id,omitempty"`

	// 子任务数量。没有子任务时为0
	SubJobsTotal *int32 `json:"sub_jobs_total,omitempty"`

	// 每个子任务的执行信息。没有子任务时为空列表
	SubJobs *[]SubJobs `json:"sub_jobs,omitempty"`
}

func (o Entities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Entities struct{}"
	}

	return strings.Join([]string{"Entities", string(data)}, " ")
}
