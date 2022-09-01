package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ObjectCompareResultDetails struct {

	// 源库名称。
	SourceDbName string `json:"source_db_name" xml:"source_db_name"`

	// 目标库名称。
	TargetDbName string `json:"target_db_name" xml:"target_db_name"`

	// 在源库的值。
	SourceDbValue *string `json:"source_db_value,omitempty" xml:"source_db_value"`

	// 在目标库的值。
	TargetDbValue *string `json:"target_db_value,omitempty" xml:"target_db_value"`

	// 错误信息。
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message"`
}

func (o ObjectCompareResultDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectCompareResultDetails struct{}"
	}

	return strings.Join([]string{"ObjectCompareResultDetails", string(data)}, " ")
}
