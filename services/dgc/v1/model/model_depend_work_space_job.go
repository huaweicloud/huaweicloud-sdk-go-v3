package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DependWorkSpaceJob struct {

	// 是否依赖最近一个周期
	DependOnLastPeriod *bool `json:"dependOnLastPeriod,omitempty"`

	// 作业名
	JobName *string `json:"jobName,omitempty"`

	// 工作空间名ID
	WorkSpaceId *string `json:"workSpaceId,omitempty"`

	// 工作空间名
	WorkSpace *string `json:"workSpace,omitempty"`

	// 依赖的规则
	DependType *string `json:"dependType,omitempty"`

	// 依赖上N个周期
	DependOnLastNPeriodValue *int32 `json:"dependOnLastNPeriodValue,omitempty"`
}

func (o DependWorkSpaceJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DependWorkSpaceJob struct{}"
	}

	return strings.Join([]string{"DependWorkSpaceJob", string(data)}, " ")
}
