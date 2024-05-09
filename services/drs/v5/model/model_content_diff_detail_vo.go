package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContentDiffDetailVo 内容对比详情信息体。
type ContentDiffDetailVo struct {

	// 源库KEY值。
	SourceKeyValue *[]string `json:"source_key_value,omitempty"`

	// 目标库KEY值。
	TargetKeyValue *[]string `json:"target_key_value,omitempty"`

	// 查询源库SQL。
	SourceSelectSql *string `json:"source_select_sql,omitempty"`

	// 查询目标库SQL。
	TargetSelectSql *string `json:"target_select_sql,omitempty"`
}

func (o ContentDiffDetailVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentDiffDetailVo struct{}"
	}

	return strings.Join([]string{"ContentDiffDetailVo", string(data)}, " ")
}
