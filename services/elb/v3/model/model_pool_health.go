package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolHealth 参数解释：后端全下线转发配置。
type PoolHealth struct {

	// 参数解释：当健康检查在线的member个数小于该个数，判定pool的state为不健康。  取值范围： - 0：默认值，不生效。 - 1：全下线转发生效。
	MinimumHealthyMemberCount *int32 `json:"minimum_healthy_member_count,omitempty"`
}

func (o PoolHealth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolHealth struct{}"
	}

	return strings.Join([]string{"PoolHealth", string(data)}, " ")
}
