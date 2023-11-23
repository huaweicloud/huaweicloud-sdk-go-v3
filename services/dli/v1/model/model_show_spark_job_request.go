package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkJobRequest Request Object
type ShowSparkJobRequest struct {

	// 批处理作业的ID。
	BatchId string `json:"batch_id"`
}

func (o ShowSparkJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkJobRequest struct{}"
	}

	return strings.Join([]string{"ShowSparkJobRequest", string(data)}, " ")
}
