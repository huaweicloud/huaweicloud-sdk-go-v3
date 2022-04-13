package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ObjectCompareResultDetails struct {
	// 源库名称。

	SourceDbName string `json:"source_db_name"`
	// 目标库名称。

	TargetDbName string `json:"target_db_name"`
	// 在源库的值。

	SourceDbValue *string `json:"source_db_value,omitempty"`
	// 在目标库的值。

	TargetDbValue *string `json:"target_db_value,omitempty"`
	// 错误信息。

	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o ObjectCompareResultDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectCompareResultDetails struct{}"
	}

	return strings.Join([]string{"ObjectCompareResultDetails", string(data)}, " ")
}
