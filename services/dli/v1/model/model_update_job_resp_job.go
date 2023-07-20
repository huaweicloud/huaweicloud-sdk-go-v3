package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateJobRespJob 作业更新信息。
type UpdateJobRespJob struct {

	// 作业更新时间, 毫秒数。
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o UpdateJobRespJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobRespJob struct{}"
	}

	return strings.Join([]string{"UpdateJobRespJob", string(data)}, " ")
}
