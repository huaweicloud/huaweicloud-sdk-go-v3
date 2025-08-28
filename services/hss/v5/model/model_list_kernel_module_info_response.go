package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKernelModuleInfoResponse Response Object
type ListKernelModuleInfoResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 内核模块列表 **取值范围**: 最小值0，最大值10000
	DataList       *[]KernelModuleInfo `json:"data_list,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListKernelModuleInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKernelModuleInfoResponse struct{}"
	}

	return strings.Join([]string{"ListKernelModuleInfoResponse", string(data)}, " ")
}
