package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProcessesResponse Response Object
type ListProcessesResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释** 进程信息列表 **取值范围** 数组长度0-10000
	DataList       *[]ProcessResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListProcessesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProcessesResponse struct{}"
	}

	return strings.Join([]string{"ListProcessesResponse", string(data)}, " ")
}
