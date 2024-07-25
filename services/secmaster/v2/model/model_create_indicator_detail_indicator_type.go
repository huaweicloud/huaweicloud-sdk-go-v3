package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIndicatorDetailIndicatorType 威胁情报类型
type CreateIndicatorDetailIndicatorType struct {

	// 威胁情报类型
	IndicatorType string `json:"indicator_type"`

	// 情报类型ID
	Id string `json:"id"`
}

func (o CreateIndicatorDetailIndicatorType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIndicatorDetailIndicatorType struct{}"
	}

	return strings.Join([]string{"CreateIndicatorDetailIndicatorType", string(data)}, " ")
}
