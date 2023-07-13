package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaurusRestartInstanceRequest 重启实例
type TaurusRestartInstanceRequest struct {

	// 实例是否延迟重启，默认false，立即重启。  - true: 延迟重启，实例将在运维时间窗内自动重启。 - false: 立即重启。
	Delay *bool `json:"delay,omitempty"`
}

func (o TaurusRestartInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaurusRestartInstanceRequest struct{}"
	}

	return strings.Join([]string{"TaurusRestartInstanceRequest", string(data)}, " ")
}
