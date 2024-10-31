package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LogConfigDto struct {

	// 防火墙id，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	FwInstanceId string `json:"fw_instance_id"`

	// 是否开启LTS，1表示是，0表示不是
	LtsEnable int32 `json:"lts_enable"`

	// LTS日志分组id,可通过查询LTS（云日志服务）下查询账号下所有日志组接口获得，通过返回值中的log_groups.log_group_id（.表示各对象之间层级的区分）获得
	LtsLogGroupId string `json:"lts_log_group_id"`

	// 攻击日志流id,可通过查询LTS（云日志服务）下查询指定日志组下的所有日志流接口获得，通过返回值中的log_streams.log_stream_id（.表示各对象之间层级的区分）获得
	LtsAttackLogStreamId *string `json:"lts_attack_log_stream_id,omitempty"`

	// 是否开启攻击日志流，1表示是，0表示不是
	LtsAttackLogStreamEnable *int32 `json:"lts_attack_log_stream_enable,omitempty"`

	// 访问控制日志流id,可通过查询LTS（云日志服务）下查询指定日志组下的所有日志流接口获得，通过返回值中的log_streams.log_stream_id（.表示各对象之间层级的区分）获得
	LtsAccessLogStreamId *string `json:"lts_access_log_stream_id,omitempty"`

	// 是否开启访问控制流，1表示是，0表示不是
	LtsAccessLogStreamEnable *int32 `json:"lts_access_log_stream_enable,omitempty"`

	// 流量日志id,可通过查询LTS（云日志服务）下查询指定日志组下的所有日志流接口获得，通过返回值中的log_streams.log_stream_id（.表示各对象之间层级的区分）获得
	LtsFlowLogStreamId *string `json:"lts_flow_log_stream_id,omitempty"`

	// 是否开启流量日志，1表示是，0表示不是
	LtsFlowLogStreamEnable *int32 `json:"lts_flow_log_stream_enable,omitempty"`
}

func (o LogConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogConfigDto struct{}"
	}

	return strings.Join([]string{"LogConfigDto", string(data)}, " ")
}
