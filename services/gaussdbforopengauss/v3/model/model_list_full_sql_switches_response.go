package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFullSqlSwitchesResponse Response Object
type ListFullSqlSwitchesResponse struct {

	// **参数解释**: 总记录数量。 **取值范围**: 不涉及。
	TotalCount *int64 `json:"total_count,omitempty"`

	// **参数解释**: 开关记录列表。
	FullSqlSwitchs *[]FullSqlSwitchResult `json:"full_sql_switchs,omitempty"`

	// **参数解释**: 可选择的SQL采集类别清单列表。供开启全量SQL时做配置下发参考。
	AllowedSqlTypes *[]SqlTypeRangeConfigResult `json:"allowed_sql_types,omitempty"`
	HttpStatusCode  int                         `json:"-"`
}

func (o ListFullSqlSwitchesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFullSqlSwitchesResponse struct{}"
	}

	return strings.Join([]string{"ListFullSqlSwitchesResponse", string(data)}, " ")
}
