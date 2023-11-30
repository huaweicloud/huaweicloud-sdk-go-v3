package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompoundMetricVoSearchResultDataValue struct {

	// CompoundMetricVO数组
	Records *[]CompoundMetricVo `json:"records,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`
}

func (o CompoundMetricVoSearchResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompoundMetricVoSearchResultDataValue struct{}"
	}

	return strings.Join([]string{"CompoundMetricVoSearchResultDataValue", string(data)}, " ")
}
