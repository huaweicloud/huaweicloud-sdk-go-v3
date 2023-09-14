package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SupplementDataResp struct {

	// 作业的开始日期 13位时间戳
	EndDate *int64 `json:"endDate,omitempty"`

	// 补数据作业名称，可能有依赖的作业，故会有多个作业的情况。
	JobList *[]string `json:"jobList,omitempty"`

	// 补数据名称
	Name *string `json:"name,omitempty"`

	// 并行周期数，取值范围[1,5]
	Parallel *int32 `json:"parallel,omitempty"`

	// 作业的结束日期 13位时间戳
	StartDate *int64 `json:"startDate,omitempty"`

	// 实例状态：SUCCESS：成功RUNNING ：运行中CANCLE：取消
	Status *string `json:"status,omitempty"`

	// 作业提交时间，13位时间戳
	SubmittedDate *int64 `json:"submittedDate,omitempty"`

	SupplementDataInstanceTime *SupplementDataRespSupplementDataInstanceTime `json:"supplement_data_instance_time,omitempty"`

	SupplementDataRunTime *SupplementDataRespSupplementDataRunTime `json:"supplement_data_run_time,omitempty"`

	// 触发补数据的类型，取值范围[0, 1]。0代表作业监控界面触发的补数据，1代表恢复动作触发的补数据
	Type *int32 `json:"type,omitempty"`

	// 用户名称
	UserName *string `json:"userName,omitempty"`
}

func (o SupplementDataResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupplementDataResp struct{}"
	}

	return strings.Join([]string{"SupplementDataResp", string(data)}, " ")
}
