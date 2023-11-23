package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkJobStatusRequest Request Object
type ShowSparkJobStatusRequest struct {

	// 批处理作业的ID。
	BatchId string `json:"batch_id"`
}

func (o ShowSparkJobStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkJobStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowSparkJobStatusRequest", string(data)}, " ")
}
