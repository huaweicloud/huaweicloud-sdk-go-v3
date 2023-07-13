package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StandardResponseResult 返回的具体结果信息
type StandardResponseResult struct {

	// 符合条件的结果列表
	Data *[]ReleaseFileVersionDo `json:"data,omitempty"`

	// 符合条件的结果总数
	TotalRecords *int32 `json:"total_records,omitempty"`

	// 符合条件的结果总页数
	TotalPages *int32 `json:"total_pages,omitempty"`
}

func (o StandardResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StandardResponseResult struct{}"
	}

	return strings.Join([]string{"StandardResponseResult", string(data)}, " ")
}
