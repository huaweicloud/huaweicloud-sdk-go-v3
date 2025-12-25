package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApiResponseData 数据信息
type ApiResponseData struct {

	// 结果数据
	Data []ApiResponseDataItem `json:"data"`

	// 每页显示的记录数
	Limit int32 `json:"limit"`

	// 当前页的起始偏移量
	Offset int32 `json:"offset"`

	// 总记录数
	Total int32 `json:"total"`
}

func (o ApiResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiResponseData struct{}"
	}

	return strings.Join([]string{"ApiResponseData", string(data)}, " ")
}
