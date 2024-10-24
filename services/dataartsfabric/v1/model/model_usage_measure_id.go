package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UsageMeasureId 使用量单位标识
type UsageMeasureId struct {
}

func (o UsageMeasureId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UsageMeasureId struct{}"
	}

	return strings.Join([]string{"UsageMeasureId", string(data)}, " ")
}
