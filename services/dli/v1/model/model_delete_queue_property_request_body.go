package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteQueuePropertyRequestBody 待删除属性
type DeleteQueuePropertyRequestBody struct {

	// 待删除队列属性key值。 范围如下： computeEngine.maxInstance computeEngine.maxPrefetchInstance job.maxConcurrent
	Keys []string `json:"keys"`
}

func (o DeleteQueuePropertyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteQueuePropertyRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteQueuePropertyRequestBody", string(data)}, " ")
}
