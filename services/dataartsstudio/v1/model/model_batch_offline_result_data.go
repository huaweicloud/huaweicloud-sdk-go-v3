package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchOfflineResultData 批量下线的返回结果，成功的个数。
type BatchOfflineResultData struct {
	Value *BatchOperationVo `json:"value,omitempty"`
}

func (o BatchOfflineResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOfflineResultData struct{}"
	}

	return strings.Join([]string{"BatchOfflineResultData", string(data)}, " ")
}
