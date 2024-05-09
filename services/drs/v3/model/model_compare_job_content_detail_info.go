package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompareJobContentDetailInfo 内容对比详情信息体。
type CompareJobContentDetailInfo struct {

	// 源库KEY值列表。
	SourceKeyValue *[]string `json:"source_key_value,omitempty"`

	// 目标库KEY值列表。
	TargetKeyValue *[]string `json:"target_key_value,omitempty"`

	// 查询源库的SQL。
	SelectSql *string `json:"select_sql,omitempty"`

	// 查询目标库的SQL。
	TargetSelectSql *string `json:"target_select_sql,omitempty"`
}

func (o CompareJobContentDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareJobContentDetailInfo struct{}"
	}

	return strings.Join([]string{"CompareJobContentDetailInfo", string(data)}, " ")
}
