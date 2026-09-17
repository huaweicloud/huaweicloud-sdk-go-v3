package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopSlowLogInfo TopSlowLogInfo对象
type TopSlowLogInfo struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 慢SQL数量
	SlowLogNum *int64 `json:"slow_log_num,omitempty"`
}

func (o TopSlowLogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopSlowLogInfo struct{}"
	}

	return strings.Join([]string{"TopSlowLogInfo", string(data)}, " ")
}
