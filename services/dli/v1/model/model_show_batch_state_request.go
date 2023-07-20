package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchStateRequest Request Object
type ShowBatchStateRequest struct {

	// 批处理作业的ID。
	BatchId string `json:"batch_id"`
}

func (o ShowBatchStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchStateRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchStateRequest", string(data)}, " ")
}
