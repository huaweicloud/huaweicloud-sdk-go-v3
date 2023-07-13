package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetDataTransformationResp 数据加工行过滤配置信息
type GetDataTransformationResp struct {

	// 数据加工总数。
	TotalCount *int64 `json:"total_count,omitempty"`

	// 数据过滤配置信息。
	FilterConditions *[]DataTransformationInfo `json:"filter_conditions,omitempty"`
}

func (o GetDataTransformationResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDataTransformationResp struct{}"
	}

	return strings.Join([]string{"GetDataTransformationResp", string(data)}, " ")
}
