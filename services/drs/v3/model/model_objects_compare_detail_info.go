package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObjectsCompareDetailInfo 对象级对比详情信息体。
type ObjectsCompareDetailInfo struct {

	// 源库名称。
	SourceDbName *string `json:"source_db_name,omitempty"`

	// 目标库名称。
	TargetDbName *string `json:"target_db_name,omitempty"`

	// 在源库的值。
	SourceDbValue *string `json:"source_db_value,omitempty"`

	// 在目标库的值。
	TargetDbValue *string `json:"target_db_value,omitempty"`

	// 对比结果，0为不一致，2为一致，3为未完成。
	Status *int32 `json:"status,omitempty"`

	// 错误信息。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o ObjectsCompareDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectsCompareDetailInfo struct{}"
	}

	return strings.Join([]string{"ObjectsCompareDetailInfo", string(data)}, " ")
}
