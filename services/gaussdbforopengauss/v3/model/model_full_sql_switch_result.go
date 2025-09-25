package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FullSqlSwitchResult **参数解释**: 开关记录列表。 **取值范围**: 不涉及。
type FullSqlSwitchResult struct {

	// **参数解释**: 是否开启全量SQL。 **取值范围**: - true：已开启。 - false：已关闭。
	IsOpen *bool `json:"is_open,omitempty"`

	// **参数解释**: 开关状态持续的开始时间。格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量。 **取值范围**: 不涉及。
	BeginTime *string `json:"begin_time,omitempty"`

	// **参数解释**: 开关状态持续的结束时间。格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量。 **取值范围**: 为null则表示，开关状态还在持续，没有发生切换。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**: 已采集的全量SQL数据的最大保留天数。 **取值范围**: [1,30]
	SaveDays *int32 `json:"save_days,omitempty"`

	// **参数解释**: 全量SQL数据存储类型。 **取值范围**: - LTS：LTS云日志服务。
	StorageMode *string `json:"storage_mode,omitempty"`

	LtsConfig *LtsConfigInfoResult `json:"lts_config,omitempty"`

	// **参数解释**: SQL采集类型列表。 **取值范围**: 不涉及。
	SqlTypeRange *[]SqlTypeInfoResult `json:"sql_type_range,omitempty"`

	// **参数解释**: 实例ID。 **取值范围**: 不涉及。是否过滤系统用户。使能该选项后，全量SQL采集将会跳过系统用户所执行的SQL记录
	IsExcludeSysUser *bool `json:"is_exclude_sys_user,omitempty"`
}

func (o FullSqlSwitchResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullSqlSwitchResult struct{}"
	}

	return strings.Join([]string{"FullSqlSwitchResult", string(data)}, " ")
}
