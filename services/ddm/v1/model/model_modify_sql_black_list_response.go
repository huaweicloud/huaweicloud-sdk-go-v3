package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifySqlBlackListResponse Response Object
type ModifySqlBlackListResponse struct {

	// **参数解释**：  全量匹配sql黑名单。  **参数范围**：  不涉及。
	SqlBlackListFullCheck *string `json:"sql_black_list_full_check,omitempty"`

	// **参数解释**：  前缀匹配sql黑名单。  **参数范围**：  不涉及。
	SqlBlackListPrefixCheck *string `json:"sql_black_list_prefix_check,omitempty"`

	// **参数解释**：  正则匹配sql黑名单。  **参数范围**：  不涉及。
	SqlBlackListRegexCheck *string `json:"sql_black_list_regex_check,omitempty"`
	HttpStatusCode         int     `json:"-"`
}

func (o ModifySqlBlackListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifySqlBlackListResponse struct{}"
	}

	return strings.Join([]string{"ModifySqlBlackListResponse", string(data)}, " ")
}
