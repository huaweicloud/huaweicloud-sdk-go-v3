package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobBatchDelete struct {

	// 作业ID列表。获取方法，请参见[获取作业ID](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)。
	JobIdList *[]string `json:"job_id_list,omitempty"`
}

func (o JobBatchDelete) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobBatchDelete struct{}"
	}

	return strings.Join([]string{"JobBatchDelete", string(data)}, " ")
}
