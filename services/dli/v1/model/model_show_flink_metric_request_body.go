package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlinkMetricRequestBody 查询作业监控信息
type ShowFlinkMetricRequestBody struct {

	// 作业ID列表。
	JobIds []int64 `json:"job_ids"`
}

func (o ShowFlinkMetricRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkMetricRequestBody struct{}"
	}

	return strings.Join([]string{"ShowFlinkMetricRequestBody", string(data)}, " ")
}
