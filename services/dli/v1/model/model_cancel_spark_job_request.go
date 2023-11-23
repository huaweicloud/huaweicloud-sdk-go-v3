package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelSparkJobRequest Request Object
type CancelSparkJobRequest struct {

	// 批处理作业的ID。
	BatchId string `json:"batch_id"`
}

func (o CancelSparkJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSparkJobRequest struct{}"
	}

	return strings.Join([]string{"CancelSparkJobRequest", string(data)}, " ")
}
