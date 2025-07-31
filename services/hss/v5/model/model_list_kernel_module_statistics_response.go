package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKernelModuleStatisticsResponse Response Object
type ListKernelModuleStatisticsResponse struct {

	// 内核模块统计信息总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 内核模块统计信息列表
	DataList       *[]KernelModuleStatisticsResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ListKernelModuleStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKernelModuleStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListKernelModuleStatisticsResponse", string(data)}, " ")
}
