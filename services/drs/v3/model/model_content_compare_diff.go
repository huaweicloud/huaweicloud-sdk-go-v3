package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ContentCompareDiff struct {

	// 查询目标库的SQL。
	TargetSelectSql *string `json:"target_select_sql,omitempty" xml:"target_select_sql"`

	// 查询源库的SQL。
	SourceSelectSql *string `json:"source_select_sql,omitempty" xml:"source_select_sql"`

	// 源库KEY值列表。
	SourceKeyValue []string `json:"source_key_value" xml:"source_key_value"`

	// 目标库KEY值列表。
	TargetKeyValue []string `json:"target_key_value" xml:"target_key_value"`
}

func (o ContentCompareDiff) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentCompareDiff struct{}"
	}

	return strings.Join([]string{"ContentCompareDiff", string(data)}, " ")
}
