package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateQueueRequestBodyProperties 添加队列属性。
type CreateQueueRequestBodyProperties struct {

	// 是否使用DLI Native。当前只涉及开启两种算子：Scan 和 Filter。修改现有队列的本属性，需要重启队列才会生效。
	ComputeEngineSparkNativeEnabled *string `json:"computeEngine.spark.nativeEnabled,omitempty"`
}

func (o CreateQueueRequestBodyProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQueueRequestBodyProperties struct{}"
	}

	return strings.Join([]string{"CreateQueueRequestBodyProperties", string(data)}, " ")
}
