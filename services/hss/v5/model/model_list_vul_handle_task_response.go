package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulHandleTaskResponse Response Object
type ListVulHandleTaskResponse struct {

	// **参数解释**: 记录总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 扫描任务列表 **取值范围**: 最小值0，最大值200
	DataList       *[]VulHandleTaskResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListVulHandleTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulHandleTaskResponse struct{}"
	}

	return strings.Join([]string{"ListVulHandleTaskResponse", string(data)}, " ")
}
