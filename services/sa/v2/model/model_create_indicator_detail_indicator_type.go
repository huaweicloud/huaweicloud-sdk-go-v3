package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 指标类型统计数据
type CreateIndicatorDetailIndicatorType struct {

	// indicator_type
	IndicatorType *string `json:"indicator_type,omitempty"`

	// id
	Id *string `json:"id,omitempty"`

	// category
	Category *string `json:"category,omitempty"`

	// layoutId
	LayoutId *string `json:"layoutId,omitempty"`
}

func (o CreateIndicatorDetailIndicatorType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIndicatorDetailIndicatorType struct{}"
	}

	return strings.Join([]string{"CreateIndicatorDetailIndicatorType", string(data)}, " ")
}
