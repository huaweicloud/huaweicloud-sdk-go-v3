package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePerformanceResourceReq struct {

	// 运行的最大作业数量
	JobQuota int32 `json:"job_quota"`
}

func (o UpdatePerformanceResourceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePerformanceResourceReq struct{}"
	}

	return strings.Join([]string{"UpdatePerformanceResourceReq", string(data)}, " ")
}
