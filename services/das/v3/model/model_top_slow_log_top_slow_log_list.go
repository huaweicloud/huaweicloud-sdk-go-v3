package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopSlowLogTopSlowLogList struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例数量
	InstanceName *string `json:"instance_name,omitempty"`

	// 慢SQL数量
	SlowLogNum *int64 `json:"slow_log_num,omitempty"`
}

func (o TopSlowLogTopSlowLogList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopSlowLogTopSlowLogList struct{}"
	}

	return strings.Join([]string{"TopSlowLogTopSlowLogList", string(data)}, " ")
}
