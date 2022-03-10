package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新云日志请求参数。
type UpdateLogtankOption struct {
	// 日志组别id，其他（非ELB）服务提供

	LogGroupId *string `json:"log_group_id,omitempty"`
	// 日志订阅主题id，其他（非ELB）服务提供

	LogTopicId *string `json:"log_topic_id,omitempty"`
}

func (o UpdateLogtankOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogtankOption struct{}"
	}

	return strings.Join([]string{"UpdateLogtankOption", string(data)}, " ")
}
