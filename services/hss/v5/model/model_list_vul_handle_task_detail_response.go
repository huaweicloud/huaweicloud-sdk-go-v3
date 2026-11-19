package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulHandleTaskDetailResponse Response Object
type ListVulHandleTaskDetailResponse struct {

	// **参数解释**: 记录总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 处置任务的详情列表 **取值范围**: 最小值0，最大值200
	DataList       *[]VulHandleTaskDetailResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListVulHandleTaskDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulHandleTaskDetailResponse struct{}"
	}

	return strings.Join([]string{"ListVulHandleTaskDetailResponse", string(data)}, " ")
}
