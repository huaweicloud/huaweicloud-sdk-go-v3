package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelFlinkJobRequestBody **参数解释**：取消Flink作业的请求体。 **约束限制**：不涉及。
type CancelFlinkJobRequestBody struct {
	CancelStrategy *FlinkJobCancelStrategy `json:"cancel_strategy"`
}

func (o CancelFlinkJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelFlinkJobRequestBody struct{}"
	}

	return strings.Join([]string{"CancelFlinkJobRequestBody", string(data)}, " ")
}
