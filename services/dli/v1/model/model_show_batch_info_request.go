package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchInfoRequest Request Object
type ShowBatchInfoRequest struct {

	// 批处理作业的ID。
	BatchId string `json:"batch_id"`
}

func (o ShowBatchInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchInfoRequest", string(data)}, " ")
}
