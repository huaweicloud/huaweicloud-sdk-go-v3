package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteSubJobsReq struct {

	// 任务ID列表，只能删除SUCCESS、FAILED两种状态。job_ids和delete_by_status必传一个。
	JobIds *[]string `json:"job_ids,omitempty"`

	// 通过任务状态删除，只能删除以下的两种状态 SUCCESS：成功。 FAILED：失败。job_ids和delete_by_status必传一个。
	DeleteByStatus *string `json:"delete_by_status,omitempty"`
}

func (o DeleteSubJobsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubJobsReq struct{}"
	}

	return strings.Join([]string{"DeleteSubJobsReq", string(data)}, " ")
}
