package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LogConfigDto struct {

	// 防火墙id
	FwInstanceId string `json:"fw_instance_id"`

	// 是否开启LTS
	LtsEnable int32 `json:"lts_enable"`

	// LTS日志分组id
	LtsLogGroupId string `json:"lts_log_group_id"`

	// 攻击日志流id
	LtsAttackLogStreamId *string `json:"lts_attack_log_stream_id,omitempty"`

	// 是否开启攻击日志流
	LtsAttackLogStreamEnable int32 `json:"lts_attack_log_stream_enable"`

	// 访问控制日志流id
	LtsAccessLogStreamId *string `json:"lts_access_log_stream_id,omitempty"`

	// 是否开启访问控制流
	LtsAccessLogStreamEnable int32 `json:"lts_access_log_stream_enable"`

	// 流量日志id
	LtsFlowLogStreamId *string `json:"lts_flow_log_stream_id,omitempty"`

	// 是否开启流量日志
	LtsFlowLogStreamEnable int32 `json:"lts_flow_log_stream_enable"`
}

func (o LogConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogConfigDto struct{}"
	}

	return strings.Join([]string{"LogConfigDto", string(data)}, " ")
}
