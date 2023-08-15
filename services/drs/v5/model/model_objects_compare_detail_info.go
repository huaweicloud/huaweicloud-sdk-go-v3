package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObjectsCompareDetailInfo 对象级对比详情信息体。
type ObjectsCompareDetailInfo struct {

	// 源库对比值。
	SourceDbValue *string `json:"source_db_value,omitempty"`

	// 目标库对比值。
	TargetDbValue *string `json:"target_db_value,omitempty"`

	// 源库名称。
	SourceDbName *string `json:"source_db_name,omitempty"`

	// 目标库名称。
	TargetDbName *string `json:"target_db_name,omitempty"`

	// 失败原因。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o ObjectsCompareDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectsCompareDetailInfo struct{}"
	}

	return strings.Join([]string{"ObjectsCompareDetailInfo", string(data)}, " ")
}
