package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InitializeStandardTemplateResultData data，统一的返回结果的最外层数据结构。
type InitializeStandardTemplateResultData struct {

	// 数据标准模板字段详情数组。
	Value *[]StandElementFieldVo `json:"value,omitempty"`
}

func (o InitializeStandardTemplateResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InitializeStandardTemplateResultData struct{}"
	}

	return strings.Join([]string{"InitializeStandardTemplateResultData", string(data)}, " ")
}
