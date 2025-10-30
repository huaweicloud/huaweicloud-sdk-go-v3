package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Results struct {

	// **参数解释：** 企业项目ID。 **取值范围：** 不涉及。
	EpsId string `json:"eps_id"`

	// **参数解释：** 日志组名称。 **取值范围：** 不涉及。
	LogGroupName string `json:"log_group_name"`

	// **参数解释：** 日志组别名。 **取值范围：** 不涉及。
	LogGroupNameAlias string `json:"log_group_name_alias"`

	// **参数解释：** 日志流名称。 **取值范围：** 不涉及。
	LogStreamName string `json:"log_stream_name"`

	// **参数解释：** 日志流别名。 **取值范围：** 不涉及。
	LogStreamNameAlias string `json:"log_stream_name_alias"`

	// **参数解释：** 告警统计周期，例如：1小时。 **取值范围：** 不涉及。
	Time string `json:"time"`
}

func (o Results) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Results struct{}"
	}

	return strings.Join([]string{"Results", string(data)}, " ")
}
