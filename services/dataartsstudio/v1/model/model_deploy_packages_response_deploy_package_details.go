package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeployPackagesResponseDeployPackageDetails struct {

	// 总的异步执行的子任务个数
	AsySubtaskNum *int32 `json:"asy_subtask_num,omitempty"`

	// 异步作业id，返回给前台轮询结果
	AsyTaskId *string `json:"asy_task_id,omitempty"`

	// 发布包ID
	PackageId *int64 `json:"package_id,omitempty"`
}

func (o DeployPackagesResponseDeployPackageDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployPackagesResponseDeployPackageDetails struct{}"
	}

	return strings.Join([]string{"DeployPackagesResponseDeployPackageDetails", string(data)}, " ")
}
