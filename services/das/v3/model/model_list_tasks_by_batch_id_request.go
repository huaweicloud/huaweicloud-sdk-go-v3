package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksByBatchIdRequest Request Object
type ListTasksByBatchIdRequest struct {

	// 批次ID
	BatchId string `json:"batch_id"`
}

func (o ListTasksByBatchIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksByBatchIdRequest struct{}"
	}

	return strings.Join([]string{"ListTasksByBatchIdRequest", string(data)}, " ")
}
