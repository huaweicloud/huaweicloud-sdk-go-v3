package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKernelModuleStatisticsResponse Response Object
type ListKernelModuleStatisticsResponse struct {

	// **参数解释** 内核模块统计信息总数 **取值范围** 最小值0，最大值300000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释** 内核模块统计信息列表 **取值范围** 最小值0，最大值300000
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
