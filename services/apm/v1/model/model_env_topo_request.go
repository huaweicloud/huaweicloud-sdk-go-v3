package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnvTopoRequest 获取组件环境拓扑数据的请求参数。
type EnvTopoRequest struct {

	// 环境id。
	TargetEnvId int64 `json:"target_env_id"`

	// 方向，可为空。
	Direction *string `json:"direction,omitempty"`

	// 结束时间。
	EndTime string `json:"end_time"`

	// 开始时间。
	StartTime string `json:"start_time"`

	// 过滤。
	FilterUser *bool `json:"filter_user,omitempty"`
}

func (o EnvTopoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvTopoRequest struct{}"
	}

	return strings.Join([]string{"EnvTopoRequest", string(data)}, " ")
}
