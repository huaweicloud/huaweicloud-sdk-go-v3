package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreRedisDataResponse Response Object
type RestoreRedisDataResponse struct {

	// 参数解释： 任务ID。 约束限制： 不涉及。 取值范围： 不涉及。 默认取值： 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreRedisDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreRedisDataResponse struct{}"
	}

	return strings.Join([]string{"RestoreRedisDataResponse", string(data)}, " ")
}
