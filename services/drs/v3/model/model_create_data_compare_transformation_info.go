package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataCompareTransformationInfo 数据加工信息。
type CreateDataCompareTransformationInfo struct {

	// 加工规则，值为contentConditionalFilter。
	TransformationType string `json:"transformation_type"`

	// 过滤条件，值为sql条件语句，例如id>100，长度限制256。
	Value string `json:"value"`
}

func (o CreateDataCompareTransformationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataCompareTransformationInfo struct{}"
	}

	return strings.Join([]string{"CreateDataCompareTransformationInfo", string(data)}, " ")
}
