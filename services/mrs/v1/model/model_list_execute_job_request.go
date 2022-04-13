package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListExecuteJobRequest struct {
	// 分页查询每页返回的最大作业数量。  取值范围：[1～100]

	PageSize *string `json:"page_size,omitempty"`
	// 当前查询页码。

	CurrentPage *string `json:"current_page,omitempty"`
	// 作业名称。

	JobName *string `json:"job_name,omitempty"`
	// 集群编号。

	ClusterId string `json:"cluster_id"`
	// 作业状态编码：  - -1：Terminated表示已终止的作业状态 - 2：Running表示运行中的作业状态 - 3：Completed表示已完成的作业状态 - 4：Abnormal表示异常的作业状态

	State *string `json:"state,omitempty"`
	// 作业执行对象的编号。

	Id *string `json:"id,omitempty"`
}

func (o ListExecuteJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExecuteJobRequest struct{}"
	}

	return strings.Join([]string{"ListExecuteJobRequest", string(data)}, " ")
}
