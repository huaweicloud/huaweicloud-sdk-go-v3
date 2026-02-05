package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigSqlBlacklistReqV3Vo struct {

	// **参数解释**：  全量匹配sql黑名单。  **约束限制**：  不涉及  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	SqlBlackListFullCheck *string `json:"sql_black_list_full_check,omitempty"`

	// **参数解释**：  请求ID。  **约束限制**：  不涉及  **取值范围**：  只能由英文字母、数字组成。  **默认取值**：  不涉及。
	RequestId *string `json:"request_id,omitempty"`

	// **参数解释**：  租户在某一Region下的project ID。  获取方法请参见[获取项目ID](https://support.huaweicloud.com/api-ddm/ddm_api_01_0063.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，且长度为32个字符。  **默认取值**：  不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in09，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  逻辑库名称。  **约束限制**：  不涉及  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	LogicDbName *string `json:"logic_db_name,omitempty"`

	// **参数解释**：  前缀匹配sql黑名单。  **约束限制**：  不涉及  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	SqlBlackListPrefixCheck *string `json:"sql_black_list_prefix_check,omitempty"`

	// **参数解释**：  正则匹配sql黑名单。  **约束限制**：  不涉及  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	SqlBlackListRegexCheck *string `json:"sql_black_list_regex_check,omitempty"`
}

func (o ConfigSqlBlacklistReqV3Vo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigSqlBlacklistReqV3Vo struct{}"
	}

	return strings.Join([]string{"ConfigSqlBlacklistReqV3Vo", string(data)}, " ")
}
