package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchPublishResultData 批量发布的返回结果，成功的个数。
type BatchPublishResultData struct {
	Value *BatchOperationVo `json:"value,omitempty"`
}

func (o BatchPublishResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchPublishResultData struct{}"
	}

	return strings.Join([]string{"BatchPublishResultData", string(data)}, " ")
}
