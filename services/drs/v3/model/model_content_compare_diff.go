package model

import (
	"encoding/json"

	"strings"
)

//
type ContentCompareDiff struct {
	// 查询目标库的SQL。

	TargetSelectSql *string `json:"target_select_sql,omitempty"`
	// 查询源库的SQL。

	SourceSelectSql *string `json:"source_select_sql,omitempty"`
	// 源库KEY值列表。

	SourceKeyValue []string `json:"source_key_value"`
	// 目标库KEY值列表。

	TargetKeyValue []string `json:"target_key_value"`
}

func (o ContentCompareDiff) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ContentCompareDiff struct{}"
	}

	return strings.Join([]string{"ContentCompareDiff", string(data)}, " ")
}
