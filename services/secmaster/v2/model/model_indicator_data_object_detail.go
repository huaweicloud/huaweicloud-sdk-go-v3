package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// detail of indicator
type IndicatorDataObjectDetail struct {
	IndicatorType *CreateIndicatorDetailIndicatorType `json:"indicator_type,omitempty"`

	// 值，如：ip url domain等
	Value *string `json:"value,omitempty"`
}

func (o IndicatorDataObjectDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndicatorDataObjectDetail struct{}"
	}

	return strings.Join([]string{"IndicatorDataObjectDetail", string(data)}, " ")
}
