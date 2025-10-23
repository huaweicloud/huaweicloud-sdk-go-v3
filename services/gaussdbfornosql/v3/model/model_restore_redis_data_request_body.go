package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreRedisDataRequestBody 参数解释： 恢复到已有实例的请求body。 约束限制： 不涉及。 取值范围： 不涉及。 默认取值： 不涉及。
type RestoreRedisDataRequestBody struct {
	RecoveryInfo *RecoveryInfo `json:"recovery_info,omitempty"`
}

func (o RestoreRedisDataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreRedisDataRequestBody struct{}"
	}

	return strings.Join([]string{"RestoreRedisDataRequestBody", string(data)}, " ")
}
